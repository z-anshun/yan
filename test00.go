package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)
type Food struct {
id int
pri int
}
/*学长喜欢的美食有有：

- 小面 6元
- 饭团 7元
- 香菇滑鸡 12元
- 小炒肉 15元
- 黄焖鸡 16元
- 冒菜 18元

### 题目要求

你需要做的是，把这些美食添加到你自己新建的一个美食表里

每次运行程序时，按照一定的规则（可以自行设定，如随机挑选、每两次挑选的食物价格不超过20元等等）在控制台打印出挑选出的食物和它的价格
*/
var sum int
var s []int
func main() {
	//打开数据库
	db, err:= sql.Open("mysql",
		"root:123@tcp(localhost:3306)/food?charset=utf8")
	if err!=nil{
		fmt.Println("打开数据库失败")
		return
	}
	stmt, e := db.Prepare("select * from delicious where id=?")
	db.Ping()
	if e!=nil{
		fmt.Println(e.Error() )
		return
	}
	//创建查找函数
	find(stmt )
	defer db.Close()
	defer  stmt.Close()
}
func find(stmt *sql.Stmt) {
	k:=0
	for{
		fmt.Println("您可以挑选您想吃的美食(选择编号即可)：")
		fmt.Printf("1)小面 6元\t2)饭团 7元\n")
		fmt.Printf("3)香菇滑鸡 12元\t4)小炒肉 15元\n")
		fmt.Printf("5)黄焖鸡 16元\t6)冒菜 18元\n")
		fmt.Println("7)显示已选\t8)quit")
		_, err := fmt.Scanln(&k)
		if  err!=nil||k<1||k>8{
			fmt.Println("请正确输入编号")
			continue
		}
		if k==8{
			fmt.Println("bye")
			os.Exit(0)
		} else if k==7 {
			show()
		}else{
			add(k,stmt )//返回查找的
		}
		if sum>=30{
			fmt.Println("您已选了超过30元的东西了，请谨慎考虑！！")
		}
	}
}
func add(i int,stmt *sql.Stmt ){
		row, e := stmt.Query(i)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		var food Food
		row.Next()
		row.Scan(&food.id, &food.pri)
		sum += food.pri
	s=append(s,i) //这个是为储存已选项
}
//一个一个的遍历
func show(){
	fmt.Println("已选择")
	for _,i:=range s{
		switch i{
		case 1:
			fmt.Println("小面")
		case 2:
			fmt.Println("饭团")
		case 3:
			fmt.Println("香菇滑鸡")
		case 4:
			 fmt.Println("小炒肉")
		case 5:
			fmt.Println("黄焖鸡")
		case 6:
			fmt.Println("冒菜")
		}
	}
	fmt.Println("总价格：",sum)
}
