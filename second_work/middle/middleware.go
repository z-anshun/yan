package middle

import (
	"github.com/gin-gonic/gin"

	"src/second_work/jwt"
	"src/second_work/model"
	"src/second_work/resp"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从header里取
		auth := c.GetHeader("Authorization")
		if len(auth) < 7 {
			//未找到
			c.Next()
		}
		token := auth[7:]
		var j jwt.Jwt
		if err := j.CheckJwt(token); err != nil {
			resp.TokenError(c, err.Error())
		}
		c.Set("name", j.UserName)

		c.Next()
		return
	}
}

//判断是否在最后10秒
func CheckTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if model.T == 0 {
			resp.Untime(c) //时间未到
		}
		if model.Final && model.T+10 > time.Now().Unix() { //到时间了
			c.Next()
		}
		if model.T+10 < time.Now().Unix() {
			resp.OverTime(c)
		}

	}
}
