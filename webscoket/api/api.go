package api

import (
	"github.com/gin-gonic/gin"
)

func postApi(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func SetRouter(e *gin.Engine) {
	//聊天室用的是get
	e.Use(postApi)
	//群聊
	e.GET("/ws", WsPage)
	e.POST("/login", Login)
	//单聊

	//多聊

}
