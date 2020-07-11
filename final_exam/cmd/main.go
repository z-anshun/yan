package main

import (
	"final_exam/api"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()


	api.SetRouter(engine)

	engine.Run(":8080")
}
