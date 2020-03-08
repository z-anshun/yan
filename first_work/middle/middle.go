package middle

import (
	"github.com/gin-gonic/gin"
	"src/first_work/jwt"
	"src/first_work/rsp"
)
//检查token
func MiddleWare(c *gin.Context) {

	//从header里取
	auth := c.GetHeader("Authorization")
	if len(auth) < 7 {
		c.JSON(401, gin.H{
			"code":    "008",
			"message": "no token",
		})
		c.Abort()
		return
	}
	token := auth[7:]
	uid, name, err := jwt.CheckToken(token)
	if err != nil {
		rsp.TokenError(c)
		c.Abort()
		return
	}
	c.Set("name",name)
	c.Set("uid",uid) //直接扔进去
	c.Next()
	return

}
