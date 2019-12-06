package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type User struct {
	Name     string `from:"name"`
	Id       string `json:"id",from:"id"`
	Password string `from:"password"`
	Message  string `from:"message"`
	Pid      string `from:"pid"`
	Ppid     string `from:"ppid"`
	Vip      string `from:"vip"`
}

/*var db *sql.DB //定义一个全局的db
func init() {
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}
我也不知道为毛先用这个，后面再使用db时都会panic红一片
}*/

//单纯的发布
func SendMsg(c *gin.Context) {
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	var name string
	var password string
	name, err = c.Cookie("first") //从cookie1中拿取名字
	if err != nil {
		name = c.PostForm("name")
		/*c.JSON(500, gin.H{
			"status":  http.StatusForbidden,
			"message": "Cookie 读取失败",
		})
		return*/
	}
	password, err = c.Cookie("second") //从cookie2中拿取密码
	if err != nil {
		/*c.JSON(500, gin.H{
			"status":  http.StatusForbidden,
			"message": "Cookie 读取失败",
		})
		return*/
		password = c.PostForm("password")
	}
	message := c.PostForm("message") //看它发布的信息
	if Usersignin(name, password) {  //判断是否登录正确
		stmt, err := db.Prepare(
			"insert into mes (`name`,message) values(?,?)  ") //写入留言
		if err != nil {
			fmt.Println("发表留言失败", err.Error())
			return
		}
		r, err := stmt.Query(name, message)
		if err != nil {
			fmt.Println(err.Error())
			return
		} //发表后可查看该条留言的所有
		r.Next()
		var id string
		var pid string
		r.Scan(&id, &name, &pid, &message)
		c.Writer.WriteString(id + "  " + name + "  " + message) //这里只是查看该条留言
	} else {
		c.Writer.WriteString("请先登录")
		c.Redirect(http.StatusMovedPermanently, "/logup") //重定向为登录
	}
}

//回复留言和评论均可
func Retmessage(c *gin.Context) {
	name, err := c.Cookie("first") //回复留言肯定要先登录撒
	if err != nil {
		c.JSON(500, gin.H{
			"status":  http.StatusForbidden,
			"message": "Cookie 读取失败",
		})
		name = c.PostForm("name") //cookie中没拿到就从请求中获取
		//return
	}
	password, err := c.Cookie("second")
	if err != nil {
		c.JSON(500, gin.H{
			"status":  http.StatusForbidden,
			"message": "Cookie 读取失败",
		})
		password = c.PostForm("password")
		//return
	}
	message := c.PostForm("message")
	pid := c.PostForm("pid")
	if Usersignin(name, password) {
		db, _ := sql.Open("mysql",
			"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
		db.SetMaxOpenConns(1000)
		err := db.Ping() //启动
		if err != nil {
			fmt.Println("连接数据库失败")
			return
		}
		if _, err = db.Query("select * from mes where id=" + pid); err == nil { //这里判断是否有着这条评论
			stmt, err := db.Prepare("insert into mes (name,message,pid) values(?,?,?) ")
			if err != nil {
				fmt.Println("预处理失败", err.Error())
				return
			}
			defer stmt.Close()
			r, err := stmt.Exec(name, message, pid)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			count, _ := r.RowsAffected()
			if count > 0 {
				fmt.Println("回复成功")
			} else {
				fmt.Println("回复失败")
			}
		} else {
			c.Writer.WriteString("该条评论不存在")
			return
		}
	} else {
		c.Writer.WriteString("请先登录")
		c.Redirect(http.StatusMovedPermanently, "/logup")
	}
}

//留言的查看 单纯的留言查看 可看单独的和一共
func Getmessage(c *gin.Context) {
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	id, flag := c.GetQuery("id") //从申请中获取id
	if !flag {
		fmt.Println("获取失败")
	}
	var id_m string
	var name string
	var message string
	var pid string //定义这些，，感觉结构体要舒服点
	if id == "" {
		r, err := db.Query("select * from mes  ")
		if err != nil {
			fmt.Println(err.Error())
			c.Abort()
			return
		}
		for r.Next() {
			r.Scan(&id_m, &name, &pid, &message)
			if pid == "0" { //只察看留言，不包括评论
				c.Writer.WriteString(id + "  " + name + "  " + message)
			}
		}

	} else {
		stmt, err := db.Prepare("select * from mes where id = ?")
		if err != nil {
			fmt.Println(err.Error())
			c.Abort()
			return
		}
		id_, _ := strconv.Atoi(id)
		r, err := stmt.Query(id_)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		r.Next()
		err = r.Scan(&id_m, &name, &pid, &message)
		if err != nil {
			fmt.Println(err.Error())
			c.Abort()
			return
		}
		c.Writer.WriteString(id_m + "  " + name + "  " + message)
	}

}

//察看单条留言加评论
func Lea_mes(c *gin.Context) {
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	id := c.Query("id")                                                  //获取想看到的留言及其评论
	stmt, err := db.Prepare("select * from mes where pid = ? or id = ?") //预处理
	if err != nil {
		fmt.Println(err.Error())
		c.Abort()
		return
	}
	rows, err := stmt.Query(id, id)
	if err != nil {
		fmt.Println("获取失败", err.Error())
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Pid, &user.Message) //赋值
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.Writer.WriteString(user.Id + "  " + user.Name + "  " + user.Message)
		c.Writer.WriteString("\n") //单纯为了好看
	}

}
func Userup(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	UsersingUP(name, password)
	return
}
func Userin(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	if !Usersignin(name, password) {
		c.Writer.WriteString("请从新登录")
		return
	}
	return
}

//点赞
//pid表示赞的是哪条，可以是评论，也是是留言
func Support(c *gin.Context) {
	var user User
	flag := true //flag是判断是否点过赞的
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	user.Name, err = c.Cookie("first")
	if err != nil {
		user.Name = c.PostForm("name")
	}
	user.Password, err = c.Cookie("second")
	if err != nil {
		user.Password = c.PostForm("password")
	}
	user.Pid = c.PostForm("pid")
	if Usersignin(user.Name, user.Password) {
		rows, err := db.Query("select * from support ") //我也不晓得为毛这里不能用where name=？
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for rows.Next() {
			var user_p User
			rows.Scan(&user.Id, &user_p.Name, &user_p.Pid, &user_p.Ppid) //这里是外面哪个user接的id
			if user_p.Name == user.Name && user_p.Pid == user.Pid {      //这里代表该用户点了这个赞了
				flag = false
				break
			}
		}
		if !flag { //代表没搜到该用户点过赞，则要insert
			fmt.Println(err.Error())
			stmt, err := db.Prepare("insert into support (name,pid,ppid) values (?,?,1)")
			if err != nil {
				fmt.Println("预处理失败")
				return
			}
			r, err := stmt.Exec(user.Name, user.Pid) //ppid代表点赞，1代表赞，0代表不赞
			if err != nil {                          //pid表示点哪个
				fmt.Println(err.Error())
				return
			}
			i, _ := r.RowsAffected()
			if i > 0 {
				fmt.Println("点赞成功")
			} else {
				fmt.Println("点赞失败")
			}
		} else {
			_, err := db.Prepare("update support set ppid=0 where id" + user.Id) //取消点赞
			if err != nil {
				fmt.Println("取消点赞失败：", err.Error())
				return
			} else {
				fmt.Println("取消点赞成功")
			}

		}
	} else {
		c.Writer.WriteString("请先登录")
		c.Redirect(http.StatusMovedPermanently, "/logup")
	}
}
func Delete(c *gin.Context) {

	name, err := c.Cookie("first")
	if err != nil {
		name, _ = c.GetQuery("name")
	}
	password, err := c.Cookie("second")
	if err != nil {
		password, _ = c.GetQuery("password")
	}
	id, _ := c.GetQuery("id")

	var user User
	if Usersignin(name, password) {

		db, _ := sql.Open("mysql",
			"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
		db.SetMaxOpenConns(1000)
		err := db.Ping() //启动
		if err != nil {
			fmt.Println("连接数据库失败")
			return
		}

		r1, err := db.Query("select * from user ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for r1.Next() {
			//这个是看user这个表，看是否时vip，vip可任意删除
			r1.Scan(&user.Id, &user.Name, &user.Password, &user.Vip)
			if user.Name == name {
				break
			}
		}
		r2, err := db.Query("select * from mes where id=" + id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		r2.Next()
		//这个是看删除哪个,得到要删除的那条的信息
		r2.Scan(&user.Id, &user.Name, &user.Pid, &user.Message)
		if user.Vip == "1" {
			stmt, err := db.Prepare("delete from mes where id = ?")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			res, _ := stmt.Exec(id)
			count, _ := res.RowsAffected()
			if count > 0 {
				c.Writer.WriteString("删除成功")
			} else {
				c.Writer.WriteString("删除失败")
			}
		} else {
			var loacl_user User
			flag := false
			r3, _ := db.Query("select * from mes ")
			for r3.Next() {
				//查看自己发布的，看是不是被这条评论了
				r3.Scan(&loacl_user.Id, &loacl_user.Name, &loacl_user.Pid, &loacl_user.Message)
				if loacl_user.Name == name { //拿到该用户的
					//下面判断该用户的发布的id若与删除的pid（评论的）一样或者就是自己的，那么就可以删
					if user.Pid == loacl_user.Id || user.Name == loacl_user.Name {
						flag = true
						break
					}
				}
			}
			stmt, err := db.Prepare("delete  from mes where id = ?")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//可以删除自己的和评论自己的
			//这里的pid是mes表里的，即评论对应的留言的，id是该user用户发表的留言的
			if flag {
				res, _ := stmt.Exec(id)
				count, _ := res.RowsAffected()
				if count > 0 {
					c.Writer.WriteString("删除成功")
				} else {
					c.Writer.WriteString("删除失败")
				}
			} else {
				c.Writer.WriteString("你没有权限删除他人的")
				return
			}
		}
	} else {
		c.Writer.WriteString("请先登录")
		c.Redirect(http.StatusMovedPermanently, "/logup")
	}
}
func main() {

	engine := gin.Default()
	engine.POST("/logup", Userup)                //注册
	engine.POST("/login", Midd_ware(), Userin)   //登录
	engine.GET("/megs", Getmessage)              //查看所有留言（不包括评论）
	engine.POST("/megs", Midd_ware(), SendMsg)   //发表留言
	engine.GET("/megs/ID", Getmessage)           //查看单条(不包括评论）
	engine.PUT("/megs/ID", Lea_mes)              //查看单条，包括评论
	engine.POST("/megs/ID/messages", Retmessage) //回复评论
	engine.POST("/megs/sup", Support)            //点赞
	engine.DELETE("/megs/ID", Delete)            // 删除评论
	engine.Run()
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
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
	db.SetMaxOpenConns(1000)
	err := db.Ping() //启动
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	if Checkusername(name) { //实际mysql里的name已经设unique了
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
	} else {
		fmt.Println("该用户已被注册")
	}
	return
}

//判断密码是否正确
func Usersignin(name string, password string) bool {
	db, _ := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8") //连接数据库
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

//判断用户名是否重复
func Checkusername(name string) bool {
	db, err := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/users?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	var uer_login User
	ro, err := db.Query("select * from user")
	for ro.Next() {

		ro.Scan(&uer_login.Id, &uer_login.Name)
		if uer_login.Name == name {
			return false
		}
	}

	return true

}
