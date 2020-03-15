package main

import (
	"github.com/gin-gonic/gin"
	"src/second_work/api"
	"src/second_work/timer"
)


func main(){
	e := gin.Default()

	go timer.TickerStart() //异步执行定时器
	api.SetRouter(e)
	e.Run(":8080")
}

