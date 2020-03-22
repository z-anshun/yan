package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"src04/webscoket/api"
	"src04/webscoket/serve"
)



func main() {
	e:=gin.Default()
	fmt.Println("Starting application...")

	go serve.Manager.Start()
	api.SetRouter(e)

	e.Run(":8080")

	//http.HandleFunc("/ws", api.WsPage) //创建路由
	//http.ListenAndServe(":12345", nil)
}


