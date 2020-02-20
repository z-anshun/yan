package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}
type User struct {
	LoginName string `json:"login_name"`
	Pwd       string `json:"pwd"`
}

func homeHandler(c *gin.Context) {
	cname, err1 := c.Cookie("uasername")
	sid, err2 := c.Cookie("session")
	//没登录过的
	if err1 != nil || err2 != nil {
		p := &HomePage{"as"}
		t, e := template.ParseFiles("../template/home.html") //先生成一个模板文件
		if e != nil {
			log.Println("Parsing template home.html error:", e)
			return
		}
		t.Execute(c.Writer, p) //返回前端
	}
	//登录过的
	if len(cname) != 0 && len(sid) != 0 {
		http.Redirect(c.Writer, c.Request, "/userhome", http.StatusFound) //重定向
		return
	}
}
func logHander(c *gin.Context) {
	cookie, _ := c.Cookie("login_name")




	form := c.Param("login_name")
	//ger2:= c.Request.Response.Header.Get("X-login")
	fmt.Println(cookie,"next\n","\n",c.Request.Method,form,c.PostForm("pwd")  )

}

func userHomeHandler(c *gin.Context) {
	cname, err1 := c.Cookie("user_name")
	_, err2 := c.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
		return
	}

	fname := c.PostForm("username") //前端判断后再提交过来

	var p *UserPage
	if len(cname) != 0 {
		p = &UserPage{cname} //从提交的里面读
	} else if len(fname) != 0 {
		p = &UserPage{fname} //从表单里读
	}
	t, e := template.ParseFiles("./templates/userhome.html")
	if e != nil {
		log.Println("Parsing userhome.html error:", e)
		return
	}

	t.Execute(c.Writer, p) //提交渲染
}

func apiHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		c.Writer.WriteString(string(re))
		return
	}
	res, _ := ioutil.ReadAll(c.Request.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		c.Writer.WriteString(string(re))
		return
	}
	request(apibody, c)
	defer c.Request.Body.Close() //防止body留在栈里
}

func proxyHandler(c *gin.Context) {
	u, _ := url.Parse("http://127.0.0.1:9000/")    //写进部署里更好
	proxy := httputil.NewSingleHostReverseProxy(u) //替换域名
	proxy.ServeHTTP(c.Writer, c.Request)
}
