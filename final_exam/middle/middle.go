package middle

import (
	"final_exam/jwt"
	"final_exam/rsps"
	"github.com/gin-gonic/gin"

)
//检查token
func MiddleWare(c *gin.Context) {

	//从header里取
	auth := c.GetHeader("Authorization")
	//没得token无所谓
	if len(auth) < 7 {

		c.Next()
		return
	}
	token := auth[7:]
	if len(token)==0{
		return
	}
	j:=jwt.NewJwt()
	uid, name, err := j.CheckToken(token)
	if err != nil {
		rsps.TokenError(c)
		c.Abort()
		return
	}
	c.Set("name",name)
	c.Set("uid",uid) //直接扔进去
	c.Next()
	return

}
