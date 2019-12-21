package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

/*有用户、管理员两种角色

用户可以通过签到的方式获得积分，签到获得积分的限制可以自己设置（列如每天可以签到5次，每次只能获得1积分）

用户可以使用积分兑换奖品

管理员可以查看用户的积分数，设置奖品的内容和需要积分*/
var db *sql.DB
var count int
type Person struct{
	 id int
	 name string
	 pid int
	 sto int //这个是积分
	 password int
}
func Userin(c *gin.Context ){//登录和签到
	name:=c.Param("name")
	password:=c.Param("password")
	if Usersignin(name,password){
		fmt.Println("您现在点1可以签到(其它就退出)：")
		k:=0
		fmt.Scanln(&k)
		if k==1 {
			qian(name)
		}else{
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}
func Userup(c *gin.Context ){
	name:=c.Param("name")
	password:=c.Param("password")
	UsersingUP(name ,password )
}
func qian(name string ){
	s, e := db.Prepare("update user set sto=?")
	defer db.Close()
	defer s.Close()
	if e!=nil{
		fmt.Println(e.Error() )
	}
	count++
	if count>5{
		fmt.Println("你今天已经签到过了")
		return
	}
	result, _ := s.Exec(count)
	k,_:=result.RowsAffected()
	if k>0{
		fmt.Println("签到成功")
	}else{
		fmt.Println("签到失败")
	}
}
func Swi(c *gin.Context ){
	name, e := c.Cookie("first")
	if e!=nil||name==""{
		name=c.PostForm("name")
	}
	password, e := c.Cookie("second")
	if e!=nil||password ==""{
		password =c.PostForm("password")
	}
	if Usersignin(name,password ) {
		rows, e := db.Query("select * from user where name=" + name)
		if e!=nil{
			fmt.Println(e.Error() )
			return
		}
		var p Person
		rows .Next( )
		rows .Scan(&p.id ,&p.pid,&p.name ,&p.sto,p.password  )
		if p.pid==1{//这就是管理员
		fmt.Println("你现在可以查看用户的积分：")
			var p1 Person
		r,_:=db.Query("select * from user")
			for r.Next() {
				r.Scan(&p1.id, &p1.pid, &p1.name, &p1.sto, p1.password)
				if p1.pid==0{
					c.Writer.WriteString(p1.name+string( p1.sto))
				}
			}
		set()//这里可以增删，改查
			}else{//这就是用户
				r,_:=db.Query("select * for jiang")
				for r.Next(){
					var name,per string
					r.Scan(name,per)
					c.Writer.WriteString (name+per)
				}
				c.Writer .WriteString("您现在可以选择对换的礼品")
				k:=""
				fmt.Scanln(&k)
			query, _ := db.Query("select * for jiang where" + k)
			query .Next()
			var name,per string
			query.Scan(&name,&per)
			s,_:=strconv.Atoi(per)
			if p.sto>=s {
				p.sto -= s
			}
		}
		}
	}


func set(){
	for {
		fmt.Println("1)添加\t 2）删除\n 3）查看\t 4)修改\n5)quit")
		k := 0
		fmt.Scanln(&k)
		switch k {
		case 1:
			Add()
		case 2:
			Del()
		case 3:
			Show()
		case 4:
			Update()
		case 5:
			return
		}
	}

}
func main() {
	db, e := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/peo?charset=utf8")
	if e!=nil{
		fmt.Println("打开数据库失败")
		return
	}
	db.Ping()
	engine:=gin.Default()
	engine.POST("/logup", Userup)                       //注册
	engine.POST("/login", Midd_ware(), Userin) //登录
	engine.POST("/switch",Swi) //兑换礼品
}
func Midd_ware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		cookie1 := http.Cookie{
			Name:  "first",
			Value: name,
		}
		http.SetCookie(c.Writer, &cookie1)
		cookie2 := http.Cookie{
			Name:  "second",
			Value: password,
		}
		http.SetCookie(c.Writer, &cookie2) //创建2个cookie分别储存
	}
}
func UsersingUP(name string, password string) {

		stmt, err := db.Prepare("insert into user (`name`,password) values(?,?) ") //注册
		if err != nil {
			fmt.Println("注册失败", err.Error())
			return
		}
		defer stmt.Close()
		defer db.Close()

		r, err := stmt.Exec(name, password) //把name和password弄进去
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		count, _ := r.RowsAffected()
		if count > 0 {
			fmt.Println("注册成功")
		} else {
			fmt.Println("注册失败")
		}
	return
}

//判断密码是否正确
func Usersignin(name string, password string) bool {
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return false
	}

	stmt, err := db.Prepare(
		"select * from user where `name`=? and password=?")
	if err != nil {
		fmt.Println("登录失败", err.Error())
		return false
	}
	defer stmt.Close() //记得关
	defer db.Close()
	_, err = stmt.Exec(name, password)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func Add(){
	//进行预处理
	stmt, err := db.Prepare("insert into jiang values(default ,?,?)")
	//新增，，？表示占位符
	if err != nil {
		fmt.Println("处理失败  ", err.Error())
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()//正常处理后关闭
	var name,percives string
	fmt.Println("您可以输入一个奖品和所需积分")
	fmt.Scanln(&name,&percives )
	r, err := stmt.Exec(name , percives  )//这里是任意都接受
	if err != nil {
		fmt.Println("sql增添失败", err.Error())
		return
	}
	//返回结果
	count, err := r.RowsAffected()
	if err != nil {
		fmt.Println("结果获取失败", err.Error())
	}
	if count > 0 {
		fmt.Println("新增成功！")
	} else {
		fmt.Println("新增失败!")
	}

}
//删除
func Del(){
	//以序号来处理
	stmt,err:=db.Prepare("delete from jiang where name=? ")
	if err!=nil{
		fmt.Println("预处理失败 ",err.Error() )
		return
	}
	defer stmt .Close()
	var name string
	fmt.Println("你现在可以输入你想删除的奖品：")
	fmt.Scanln(&name)
	r,err:=stmt.Exec(name )
	if err!=nil{
		fmt.Println(err.Error() )
		return
	}
	count,_:=r.RowsAffected()
	if count>0{
		fmt.Println("删除成功")
	}else{
		fmt.Println("删除失败")
	}
}
//修改
func Update(){
	//该执行的命名如下
	stmt,err:=db.Prepare("update people set pervise =?," +
		" where name=?" )
	if err!=nil{
		fmt.Println("预处处理失败",err.Error())
		return
	}
	defer stmt.Close()
	var name,p string
	fmt.Println("你可以输一个姓名，地址和你想对应的编号:")
	fmt.Scanln(&name,&p )
	r,_:=stmt.Exec(p,name  )//这里相当于传参
	count,_:=r.RowsAffected() //如果修改不变，会返回一个0，表示失败
	if count >0{
		fmt.Println("修改成功")
	}else{
		fmt.Println("修改失败")
	}
}
//查询
func Show(){
	stmt,err:=db.Query("select * from jiang")
	if err!=nil{
		fmt.Println("预处理失败")
		return
	}
	defer stmt .Close()
	for  stmt.Next(){
		var p,name string
		stmt .Scan(&name,&p)//顺序不要搞错
		fmt.Println(name,p)
	}  //这里的stmt直接变成了rows
}
