package middle

import (
	"final_exam/jwt"
	"final_exam/rsps"
	"fmt"
	"github.com/gin-gonic/gin"
)

//检查token
func MiddleWare(c *gin.Context) {

	//从header里取
	token := c.PostForm("token")
	//没得token无所谓
	if len(token) == 0 {
		c.Next()
		return
	}
	j := jwt.NewJwt()
	uid, name, err := j.CheckToken(token+"=")
	if err != nil {

		rsps.TokenError(c)
		c.Next()
		return
	}
	fmt.Println("token success")

	c.Set("name", name)
	c.Set("uid", uid) //直接扔进去
	c.Next()
	return

}
