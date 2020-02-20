package main

import (
	"github.com/gin-gonic/gin"

)

func RegisterHandler() *gin.Engine {
	e := gin.Default()
	e.GET("/", homeHandler) //主界面
	e.POST("/", homeHandler)

	e.GET("/userhome", userHomeHandler) //用户登录界面
	e.POST("/userhome", userHomeHandler)

	e.POST("/login",logHander)

	e.POST("/upload/:name", proxyHandler)


	e.POST("/api", apiHandler)
	e.StaticFile("/statics", "./template")  //静态文件 127.0.0.1：8080/statics

	return e
}

func main() {
	e := RegisterHandler()
	e.Run(":8080")
}
