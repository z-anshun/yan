package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"src04/webscoket/serve"
	"src04/webscoket/util"
)

type Users struct {
	U []*User
}

type User struct {
	Uuid       string `json:"uuid"`
	UserName   string `json:"username"`
	Receipient string `json:"receipient"`
}

var Us Users

var upgrade = websocket.Upgrader{
	//实现跨域，怎么都是true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsPage(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err!=nil {
		fmt.Println("get cookie error:",err )
	}
	name:=cookie
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("websocket connect is failed,err:", err)

	}
	var receipient string

	for _, v := range Us.U {
		if v.UserName ==name {

			receipient = v.Receipient
		}
	}
	//这里的通道不能随便关
	//defer  conn.Close()
	//随便创建一个
	//这是群聊
	client := &serve.Client{Name: name, Socket: conn, Send: make(chan []byte), Recipient: receipient}

	serve.Manager.Register <- client

	go client.Read(serve.Manager)
	go client.Write()

}
func Login(c *gin.Context) {
	var user User
	user.UserName = c.PostForm("username")

	files, err := template.ParseFiles("/usr/local/workspace/src/webscoket/template/login.html")
	if err != nil {
		log.Fatal("template html error: ", err)
	}
	err = files.Execute(c.Writer, nil)

	if len(user.UserName) == 0 {
		user.UserName = "匿名用户"
	}

	user.Uuid, _ = util.NewUUID()
	//写入
	Us.U = append(Us.U, &user)

}
