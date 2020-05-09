package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"regexp"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123@(127.0.0.1:3306)/users?charset=utf8")
	if err != nil {
		log.Panicln("open mysql error:", err)
	}
	db.Ping()
	DB = db
}

func Defend(str string) bool {
	reg := regexp.MustCompile(`or(.*?)=(.*?)`)
	s := reg.FindStringSubmatch(str)
	if len(s) < 2 {
		return true
	}
	if s[1] == s[2] {
		return false
	}
	return true
}

func Repeat(name string) bool {

	row, err := DB.Query(`SELECT name FROM USER WHERE name=?`, name)
	if err != nil {
		return false
	}
	s := ""
	row.Next()
	row.Scan(&s)
	if len(s) != 0 {
		return false
	} else {
		return true
	}

}

//增添用户
func AddUser(name string, passwd string) *Resp {
	if len(name) > 15 || len(passwd) > 10 {
		return ErrorNameOrPasswd("name or password is to long")
	}

	//防注入
	if !Defend(name) || !Defend(passwd) {
		return ErrorNameOrPasswd("name or password is error")
	}

	//查重
	if !Repeat(name) {
		return ErrorNameOrPasswd("name is repeat")
	}

	stmtAdd, err := DB.Prepare("insert into user (name,password) values(?,?)")
	if err != nil {
		log.Println("add error:", err)
		return ErrorSql("add mysql error")
	}
	_, err = stmtAdd.Exec(name, passwd)
	if err != nil {
		log.Println(err)
		return ErrorSql("add exec error")
	}

	return OK("add success")

}

//验证
func Verify(name string, passwd string) *Resp {
	if !Defend(name) || !Defend(passwd) {
		return ErrorNameOrPasswd("name or password is error")
	}
	_, err := DB.Query(`select * from user where name=? and password=?`, name, passwd)
	if err != nil || err == sql.ErrNoRows {
		return ErrorVerify("login in is error")
	}

	return OK("welcome back " + name)
}

//改正
func UpdateInfromation(code int32, in string, str string) *Resp {

	//code1为改名字，，2为改密码
	if code == 1 {
		if !Repeat(str) {
			return  ErrorSql("name is repeat")
		}
		_, err := DB.Exec(`update user set name=? where name=?`, str, in)
		if err != nil {
			return ErrorSql("update name is error")
		}
	}
	if code == 2 {
		_,err:=DB.Exec(`update user set password=? where password=?`, str, in)
		if err!=nil{
			return ErrorSql("update password is error")
		}
	}
	return OK("update success")
}
