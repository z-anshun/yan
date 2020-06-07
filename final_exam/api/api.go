package api

import (
	"final_exam/middle"
	"github.com/gin-gonic/gin"
)

//全域
func postApi(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
var RoomIds []int

func SetRouter(e *gin.Engine) {

	e.Use(postApi)
	user_log := e.Group("/user")
	{
		user_log.POST("/register", Register)              //注册
		user_log.POST("/login", middle.MiddleWare, Login) //登录
		user_log.POST("/creatroom/:id", middle.MiddleWare, CreateRoom)

	}

	e.GET("/ws/:room_id", WsPage)

}
