package api

import (
	"final_exam/middle"
	"github.com/gin-gonic/gin"
)

//全域

func AccessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		c.Next()
	}
}
var RoomIds []int

func SetRouter(e *gin.Engine) {


	e.Use(AccessJsMiddleware(),middle.MiddleWare)
	user_log := e.Group("/user")
	{
		user_log.POST("/register", Register)              //注册
		user_log.POST("/login", Login) //登录
		user_log.POST("/creatroom/:id", CreateRoom)

	}

	e.GET("/ws/:room_id", WsPage)

}
