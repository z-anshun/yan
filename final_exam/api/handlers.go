package api

import (
	"database/sql"
	"errors"
	"final_exam/db"
	"final_exam/judge"
	jwt2 "final_exam/jwt"
	"final_exam/model"
	"final_exam/rsps"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var upgrade = websocket.Upgrader{
	//实现跨域，怎么都是true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//设置token和cookie
func set(c *gin.Context, user *model.User) {
	c.SetCookie("Name", user.Name, 1000, "/", "", false, true)
	jwt := jwt2.NewJwt()
	toke := jwt.CreatJwt(user.Name, user.Id)
	//放如token
	c.Header("Authorization", toke)
}

//放注入
func defend(str string) error {
	reg := regexp.MustCompile(`.*?(or)(.*?=.*?) .*?`)
	s := reg.FindStringSubmatch(str)
	if len(s) != 0 {
		v := strings.Split(s[2], "=")
		if v[1] == v[2] {
			return errors.New("")
		}
		if len(s[1]) != 0 {
			return errors.New("")
		}
	}
	return nil
}

//登录
func Login(c *gin.Context) {

	name, exists := c.Get("name")
	//是否存在
	if exists {
		rsps.Ok(c, "Login success")
		return
	}
	name = c.PostForm("name")
	pwd := c.PostForm("password")
	fmt.Println(name, pwd)
	//防注入
	if err := defend(name.(string)); err != nil {
		rsps.WordsError(c)
	}
	if err := defend(pwd); err != nil {
		rsps.WordsError(c)
	}

	if len(name.(string)) == 0 || len(pwd) == 0 {
		rsps.RegisterError(c)
		c.Abort()
		return
	}
	user := model.User{Name: name.(string), Password: pwd}
	//查找该用户
	if err := db.GetUser(&user); err != nil || err == sql.ErrNoRows {
		rsps.LoginError(c)
		c.Abort()
		return
	}
	set(c, &user)

	rsps.Ok(c, "login success")
}

//注册
func Register(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("password")

	if len(name) == 0 || len(pwd) == 0 {
		rsps.RegisterError(c)
		c.Abort()
		return
	}

	if err := db.FindUser(name); err != nil || err == sql.ErrNoRows {
		rsps.NameRepeat(c)
		return
	}
	//防注入
	if err := defend(name); err != nil {
		rsps.WordsError(c)
	}
	if err := defend(pwd); err != nil {
		rsps.WordsError(c)
	}

	user := model.User{Name: name, Password: pwd}

	if err := db.AddUser(&user); err != nil {
		rsps.RegisterError(c)
		c.Abort()
		return
	}

	set(c, &user)
	rsps.Ok(c, "register success")
}

//开一个房间
func CreateRoom(c *gin.Context) {
	id := c.Param("id")

	k, err := strconv.Atoi(id)
	//设0就是随机
	if (err == nil || len(id) != 0)&&k!=0 {
		if _, ok := model.UserMessages[k]; ok {
			rsps.CreateRooms(c)
			return
		}
	} else if k==0 {
		for {
			//房间号设为4位数
			rand.Seed(time.Now().UnixNano())
			k = rand.Intn(8999) + 1000
			if _, ok := model.UserMessages[k]; !ok {
				break
			}
		}
	}

	RoomIds = append(RoomIds, k)
	room := model.NewUsers()

	model.UserMessages[k] = room
	c.JSON(200, gin.H{
		"Code":    "000",
		"Message": "creat room success",
		"Room":    k,
	})

}

//聊天室这些
func WsPage(c *gin.Context) {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte("connect error"))
	}
	//一定要有cookie
	cookie, err := c.Cookie("Name")
	fmt.Println(cookie,err)
	if err != nil || len(cookie) == 0 {
		cookie="as"
		//conn.WriteMessage(websocket.CloseMessage, []byte("get cookie error"))
	}
	//获取房间号

	query := c.Param("room_id")
	id, err := strconv.Atoi(query)

	if err != nil {
		return
	}
	flag := 0
	for _, v := range RoomIds {

		if v == id {
			flag = 1
			break
		}
	}
	if flag == 0 {

		conn.WriteMessage(websocket.CloseMessage, []byte("get room error"))
		return
	}

	msg := model.Message{Name: cookie, Socket: conn, RoomId: id, Exit: false}
	//如果没找打房间
	_, ok := model.UserMessages[id]
	if !ok {
		conn.WriteMessage(websocket.CloseMessage, []byte("find room error"))
	}
	s := model.UserMessages[id]
	go s.Start()
	//找到对应的房间
	//进来一个人
	if len(model.UserMessages[id].Other) == 2 {

		conn.WriteMessage(websocket.TextMessage, []byte("The number of people is more"))
	} else if len(model.UserMessages[id].Other) == 0 {
		msg.Black = true //规定先进来这个是黑子
		msg.Round = true
	} else {
		msg.Black = false
		msg.Round = false
	}
	model.UserMessages[id].Other <- 1
	//开始好了
	s.Users[&msg] = true

	s.Register <- &msg
	//开始协程

	fmt.Println("start")
	go msg.Read(s)
	go msg.WriteAndDo(s)
	//如果都退出了房间
	go func() {
		//黑棋胜利
		if judge.Winter(&model.UserMessages[id].Ch.Board) == 1 {
			for k, _ := range model.UserMessages[id].Users {
				if !k.Black {
					model.UserMessages[id].Defeat <- k

				}
			}
		} else if judge.Winter(&model.UserMessages[id].Ch.Board) == -1 {
			for k, _ := range model.UserMessages[id].Users {
				if k.Black {
					model.UserMessages[id].Defeat <- k
				}
			}
		}
	}()
	go func() {
		c := 0
		m := &model.Message{}
		for k, _ := range model.UserMessages[id].Users {
			if k.Exit && k != m {
				m = k
				c++
			}
			//这里都退出了
			if c == 2 {
				//删除该房间

				delete(model.UserMessages, id)
				db.AddRecord(model.FileRecord{Record: model.UserMessages[id].Record, RoomId: id, EndTime: time.Now().Day() + 3})

			}
		}
	}()

}
