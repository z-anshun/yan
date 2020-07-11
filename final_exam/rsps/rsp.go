package rsps

import "github.com/gin-gonic/gin"
//1开头为正确，0为错误
func Ok(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"Code":    "100",
		"Message": msg,
	})
}

//登录失败
func LoginError(c *gin.Context) {
	c.JSON(401, gin.H{
		"Code":    "001",
		"Message": "Login error",
	})
}

//注册失败
func RegisterError(c *gin.Context) {
	c.JSON(401, gin.H{
		"Code":    "002",
		"Message": "Register error",
	})
}

//token读取失败
func TokenError(c *gin.Context) {
	c.JSON(401, gin.H{
		"Code":    "003",
		"Message": "token error",
	})
}
func NameRepeat(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code":    "004",
		"Message": "your name is repeat",
	})
}
func WordsError(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code":    "005",
		"Message": "your words are sensitive",
	})
}
//创房间失败
func CreateRooms(c *gin.Context){
	c.JSON(200, gin.H{
		"Code":    "006",
		"Message": " the room is repeat",
	})
}
//获得数据失败
func GetNumberError(c *gin.Context){
	c.JSON(200, gin.H{
		"Code":    "007",
		"Message": " get number error",
	})
}

