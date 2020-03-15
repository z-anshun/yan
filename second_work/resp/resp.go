package resp

import "github.com/gin-gonic/gin"

func Ok(c *gin.Context,msg string){
	c.JSON(200,gin.H{
		"code":"000",
		"message":msg,
	})

}

func TokenError(c *gin.Context,msg string){
	c.JSON(401,gin.H{
		"code":"002",
		"message":msg,
	} )
}
func RedisError(c *gin.Context,msg string){
	c.JSON(200,gin.H{
		"code":"003",
		"message":msg,
	} )
}
func GetUserError(c *gin.Context,msg string){
	c.JSON(200,gin.H{
		"code":"004",
		"message":msg,
	})
}
func UserExist(c *gin.Context){
	c.JSON(200,gin.H{
		"code":"005",
		"message":"user exist",
	} )
}

func PasswordError(c *gin.Context){
	c.JSON(200,gin.H{
		"code":"006",
		"message":"password error",
	})
}
//超时了
func OverTime(c *gin.Context){
	c.JSON(200,gin.H{
		"code":"007",
		"message":"sorry ,time is over",
	} )
}
//时间还未到
func Untime(c *gin.Context){
	c.JSON(200,gin.H{
		"code":"008",
		"message":"not yet",
	} )
}
//堵住了，，要再次投
func CoonJam(c *gin.Context)  {
	c.JSON(200,gin.H{
		"code":"009",
		"message":"sorry ,conn is jam,please once again",
	} )
}
//已经投过票的
func VotedUser(c *gin.Context){
	c.JSON(200,gin.H{
		"code":"010",
		"message":"sorry,you have voted",
	} )
}