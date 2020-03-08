package rsp

import (
	"github.com/gin-gonic/gin"
)

func DbErr(c *gin.Context) {
	c.JSON(200, gin.H{"code": "001", "message": "db error"}) //数据库错误
}
func ReadFormErr(c *gin.Context) {
	c.JSON(200, gin.H{"code": "002", "message": "form error"}) //读取form错误
}
func CheckNoOne(c *gin.Context) {
	c.JSON(200, gin.H{"code": "003", "message": "no one"}) //没找到
}
func Ok(c *gin.Context,mes string ){
	c.JSON(200,gin.H{"code":"004","message":mes} )
}

func UserExist(c *gin.Context){
	c.JSON(200,gin.H{"code":"005","message":"user exist"} )//已存在
}
func TokenError(c *gin.Context ){
	c.JSON(401,gin.H{"code":"006","message":"token error"})
}
func UserMesErr(c *gin.Context,msg string ){
	c.JSON(401,gin.H{
		"code":"007",
		"message":msg,
	})
}
