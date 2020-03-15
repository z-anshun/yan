package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"src/second_work/db"
	"src/second_work/jwt"
	"src/second_work/model"
	"src/second_work/resp"
	"strings"
)

func userGet(c *gin.Context) model.User {
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		resp.GetUserError(c, "get user message error")
		return model.User{}
	}

	//消去空格
	u.Name = strings.TrimSpace(u.Name)
	u.Password = strings.TrimSpace(u.Password)
	return u
}

//登录
func Login(c *gin.Context) {
	n, exists := c.Get("name")

	//获取到了，就是以前登录过
	if exists {
		db.LoginTimeUpdate(n.(string))
		c.JSON(200, gin.H{
			"code":    "001",
			"message": "Welcome back",
		})
		return
	}
	//如果没获取到
	u := userGet(c)
	p := db.UserGet(u.Name, "password")
	if p != u.Password {
		resp.PasswordError(c)
	}
	db.LoginTimeUpdate(u.Name)
	resp.Ok(c, "login success"+u.Name) //登录成功
}

//注册
func Register(c *gin.Context) {
	u := userGet(c)
	if u.Name == "" {
		return
	}
	db.RedisSetIn(c, u)

	j := jwt.Jwt{}

	token := j.CreatJwt(u.Name, u.Password) //创建
	c.JSON(300, gin.H{
		"code":  "000",
		"token": token,
	})
	return
}

//参赛

func JoinMatch(c *gin.Context) {
	n, exists := c.Get("name")
	if !exists {
		c.Redirect(500, "/login")
	}
	//将参赛人员的加入 在个名字后加
	db.JoinMat(c, n.(string))

	resp.Ok(c, "join match success")

}

//退赛
func LeaveMatch(c *gin.Context) {
	n, exists := c.Get("name")
	if !exists {
		c.Redirect(500, "/login")
	} //肯定要登录了才能参赛

	db.DeleteMat(c, n.(string))

	resp.Ok(c, "leave match success")

}

//投票
func Vote(c *gin.Context) {
	n, exists := c.Get("name")
	if !exists {
		c.Redirect(500, "/login")
	} //肯定要登录了才能参赛

	obj := c.PostForm("competitor") //给谁投
	if len(obj) == 0 {
		resp.GetUserError(c, "get obj error")
	}
	db.PutVote(c, n.(string), obj)

	resp.Ok(c, "vote success")
}

//获取排行榜
func Chart(c *gin.Context) {
	competitor := db.GetChart(c)         //获取排行榜
	com, err := json.Marshal(competitor) //json化
	if err != nil {
		log.Panicln("json competitor error")
	}
	c.JSON(200, string(com))
}

//最后的投票
func FinalVote(c *gin.Context) {
	n, exists := c.Get("name")
	if !exists {
		c.Redirect(500, "/login")
	} //肯定要登录了才能参赛

	if !model.Coon.GetConn() {
		log.Println("coon is jam") //管道堵塞
		resp.CoonJam(c)
	}
	obj := c.PostForm("competitor")                     //给谁投
	go db.FinalVoteUser(c, n.(string), obj, model.Coon) //最后的投票,开协程投票

}

//刷新
func FreshChart(c *gin.Context) {
	chart := db.FreshFinalChart(c)
	com, err := json.Marshal(chart) //json化
	if err != nil {
		log.Panicln("json competitor error")
	}
	c.JSON(200, string(com))
}
func FinalChart(c *gin.Context) {
	chart := db.GetFinalChart(c) //获取最后的排行榜

	com, err := json.Marshal(chart) //json化
	if err != nil {
		log.Panicln("json competitor error")
	}
	c.JSON(200, string(com))
}
