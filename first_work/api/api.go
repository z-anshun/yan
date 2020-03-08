package api

import (
	"github.com/gin-gonic/gin"
	"src/first_work/middle"
)

func SetRouter(e *gin.Engine) {
	e.POST("/register", Register) //注册
	e.POST("/login", Login)       //登录
	e.GET("/user", middle.MiddleWare, UserMsg)
	e.POST("/update", middle.MiddleWare, UpdateMsg)

}
