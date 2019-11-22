package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name     string `form:"name"`
	password int    `form:"password"`
}

var p Person
var ps []Person

func main() {
	r := gin.Default()
	routerGoup := r.Group("/user")//创建多个方法
	routerGoup.GET("/register", Hander_register)   //注册
	routerGoup.GET("/login", Hander_longin) //登录
	r.Run()
}
func Hander_register(c *gin.Context) {
	err := c.ShouldBind(&p)
	fmt.Println(c.FullPath())
	if  err != nil {
		log.Fatal(err.Error())
		c.Abort()
		return//binding失败
	}
	for _, k := range ps {
		if k.Name == p.Name {
			c.Writer.WriteString("sorry " + k.Name + "已经被注册")
			c.Abort()
			return
		}
	} //如果已经被注册了，就gg
	c.Writer.Write([]byte("注册成功   欢迎" + p.Name ))
	ps = append(ps, p)//将注册了的加入一个结构体切片
	return


}
func Hander_longin(c *gin.Context) {
	err := c.ShouldBind(&p)
	fmt.Println(c.FullPath())
	if    err != nil {
		log.Fatal(err.Error())
		c.Abort()
		return
	}//判断binding是否成功
	for _, s := range ps {
		if s==p   {
			c.Writer.WriteString("登录成功")
			c.Writer.WriteString("   欢迎回来" + p.Name)
			c.Abort()
			return
		}
		}//判断是否登录正确

		c.Writer.WriteString("登录失败")
		c.Abort()
		return
}
