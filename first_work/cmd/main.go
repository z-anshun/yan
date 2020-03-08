package main

import (
	"github.com/gin-gonic/gin"
	_ "src/first_work"
	"src/first_work/api"
)

func main() {
	e := gin.Default()

	api.SetRouter(e)

	e.Run(":8080")
}
