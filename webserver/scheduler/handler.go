package main

import (
	"github.com/gin-gonic/gin"
	"scheduler/dbops"
)

func vidDelRecHandler(c *gin.Context ){
	vid:=c.Param("name")
	if len(vid)==0{
		sendResponse(c.Writer,400,"video id should not be empty")
		return
	}
	err:=dbops.AddVdeoDeletionRecord(vid)
	if err!=nil{
		sendResponse(c.Writer ,500,"Internal server error")
		return
	}
	sendResponse(c.Writer,200,"")
	return
}