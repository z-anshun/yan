package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"os"
)

func main() {
	flie := "students"
	f, err := os.Open(flie)
	if err != nil {
		fmt.Println("打开文件失败", err.Error())
		return
	}
	reader := bufio.NewReader(f) //读取文件
	for {
		detail, err := reader.ReadString('\n') //这里表示一行一行的读取
		if err == io.EOF {
			break
		}
		str := string(detail)         //转换为string再调用相应的函数,全转换了
		str = strings.Trim(str, "\n") //把换行符清理了
		into(str[0:10], str [10:])

	}
}

func into(id_name string, name_id string) {
	db, err := sql.Open("mysql",
		"root:123@tcp(localhost:3306)/students") //打开数据库
	//数据库名：秘码@tcp(路径和接口)/表名+类型
	db.Ping()
	defer db.Close() //记得关
	if err != nil {
		fmt.Println("打开数据库失败", err.Error())
		return
	} //判断打开成功与否
	stmt, err := db.Prepare("insert into student (id,`name` ) values(?,?)") //进行预处理
	if err != nil {
		fmt.Println("预处理失败", err.Error())
		return
	}
	defer stmt.Close()

	r, err := stmt.Exec(id_name  , name_id)//放入，执行命令
	if err != nil {
		fmt.Println("添加失败", err.Error())
		return
	}
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("增添成功")
	} else {
		fmt.Println("增添失败")
	} //判断添加成功与否
}
