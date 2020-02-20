package main

import (
	"github.com/gin-gonic/gin"
	"taskrunner"
)

func RegisterHandlers()*gin.Engine {
	e:=gin.Default()
	e.POST("/video-delete-record/:name",vidDelRecHandler)
	return  e
}

func main(){
	go taskrunner.Start()
	e:=RegisterHandlers()
	e.Run(":9001")//阻塞模式
}
