package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func midWare()gin.HandlerFunc{
	return func( c *gin.Context) {
		if !M.GetConn() {
			sendErrorResponse(c.Writer, http.StatusTooManyRequests, "Too many requests")
		}

	}
}
func RegisterHandler()*gin.Engine{
	e:=gin.Default()
	e.GET("/videos/:vid-id",midWare(),streamHander)
	e.POST("/upload/:vid-id", uploadHandler)
	e.GET("/testpage",testPageHandler)

	e.POST("/upload/av/:user_name",uploadAv)
	e.GET("/av/:user_name",getAv)
	return  e
}

func main(){
	e:=RegisterHandler()
	http.ListenAndServe(":9000",e)
}
