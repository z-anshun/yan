package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"src/first_work"
	"src/first_work/jwt"
	"src/first_work/rsp"
)

type UserForm struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var f UserForm
	j:=jwt.NewJwt()
	if err := c.ShouldBindJSON(&f); err != nil {
		rsp.ReadFormErr(c)
		return
	} //读取用户的输入

	var u User
	first_work.DB.Where("name=?", f.UserName).First(&u)
	if u.ID != 0 {
		rsp.UserExist(c)
		return
	}
	u.Name = f.UserName
	u.Password = f.Password

	first_work.DB.Create(&u)

	token := j.CreatJwt(u.Name, int(u.ID))
	
	c.JSON(200,gin.H{
		"token":token,
	} )
	return

}
func Login(c *gin.Context) {
	var f UserForm
	if err := c.ShouldBindJSON(&f); err != nil {
		rsp.ReadFormErr(c)
		return
	}
	var u User
	first_work.DB.Where(User{
		Name:     f.UserName,
		Password: f.Password,
	}).First(&u)
	if u.ID == 0 {
		rsp.CheckNoOne(c)
		return
	}
	
	rsp.Ok(c, "user login")
}

//用户信息
func UserMsg(c *gin.Context) {
	name, err := c.Get("name")

	if !err || len(name.(string)) == 0 {
		rsp.UserMesErr(c, "get user message error")
		return
	}
	var u User
	first_work.DB.Where("name=?", name).First(&u) //查询失败
	if u.ID == 0 {
		rsp.UserMesErr(c, "get user message error from db")
		return
	}
	
	c.JSON(200, u)
}

//更改用户信息
func UpdateMsg(c *gin.Context) {
	name, err := c.Get("name")

	if !err || len(name.(string)) == 0 {
		rsp.UserMesErr(c, "get user message error")
		return
	}
	key := c.PostForm("key")     //key表示用户更改的
	value := c.PostForm("value") //v表示值

	if len(key) == 0 {
		rsp.UserMesErr(c, "get key error")
		return
	}
	var u User
	if key == "name" {
		first_work.DB.Where("name=?", value).First(&u)
		if u.ID != 0 {
			rsp.UserExist(c) //看该名字是否被占用
			return
		}
	}
		e := first_work.DB.Table("users").Where("name=?", name).Update(key, value).Error
	if e != nil {
		c.JSON(401,
			gin.H{"code": "009",
				"message": "update error"})
	}

	
	rsp.Ok(c, "update success")

}
