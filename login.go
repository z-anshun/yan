package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Age      int    `form:"age"`
}

func main() {
	r := gin.Default()
	r.POST("/user/*action", Write) //创立post请求，接受用户的创立
	r.Run()
}
func Write(c *gin.Context) {
	var per User
	err := c.Bind(&per)
	if err != nil {
		log.Fatal(err.Error())
		c.Abort()
		return
	}//判断绑定
	cookie := &http.Cookie{
		Name:     per.Name,
		Value:    per.Password,
		Path:     "/",
		HttpOnly: true,
	}//来个cookie，储存用户信息
	http.SetCookie(c.Writer, cookie)//把cookie存放
	if cookie.Name != "" {
		c.JSON(200, gin.H{
			"code":    200,
			"message": cookie.Name + "hello",
		})//判断是注册，还是啥子都没敢
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "hello guest",
		})//啥子都没干
	}
}
