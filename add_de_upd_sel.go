package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", //这是一个驱动
		//用户名：密码@协议（ip地址:端口）/数据库
		"root:123@tcp(localhost:3306)/first?charset=utf8")
	db.Ping() //真正打开连接,要么一开始，要么就最后自动连接

	defer func() {
		if db != nil {
			db.Close()
		}
	}() //关闭，正常处理后，就关了
	if err != nil {
		fmt.Println("连接失败")
		return
	} //错误处理
	//来波选择处理
	var i int
	fmt.Println("-------请输入你想进行操作的编号-------")
	fmt.Println("1)增添               2）删除")
	fmt.Println("3)修改               4）查询整个表")
	fmt.Println("5)exit")
	_,err=fmt.Scan("%d",&i)
	if err!=nil{
		fmt.Println("输入有误")
		return
	}
	//选择判断
	switch i  {
	case 1:add(db) //直接把DB传过去
	case 2:delete(db)
	case 3:update(db)
	case 4:sel(db)
	default:
		 return
	}

}
//增添
	func add(db *sql.DB ){
		//进行预处理
		stmt, err := db.Prepare("insert into people value(default ,?,?)")
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
		var name,address string
		fmt.Println("您可以输入一个姓名和地址")
		fmt.Scanln(&name,&address)
	r, err := stmt.Exec(name , address )//这里是任意都接受
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
	//来波获取
	id,_:=r.LastInsertId()
	fmt.Println("共输入了 ",id,"人")

}
//删除
func delete(db*sql.DB ){
	//以序号来处理
	stmt,err:=db.Prepare("delete from people where id=? ")
	if err!=nil{
		fmt.Println("预处理失败 ",err.Error() )
		return
	}
	defer stmt .Close()
	var id int
	fmt.Println("你现在可以输入你想删除的序号：")
	fmt.Scanln(&id)
	r,err:=stmt.Exec(id)
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
func update(db *sql.DB ){
	//该执行的命名如下
	stmt,err:=db.Prepare("update people set name =?," +
		"address=? where id=?" )
	if err!=nil{
		fmt.Println("预处处理失败",err.Error())
		return
	}
	defer stmt.Close()
	var id int
	var name,address string
	fmt.Println("你可以输一个姓名，地址和你想对应的编号:")
	fmt.Scanln(&name,&address ,&id )
	r,_:=stmt.Exec(name ,address ,id )//这里相当于传参
	count,_:=r.RowsAffected() //如果修改不变，会返回一个0，表示失败
	if count >0{
		fmt.Println("修改成功")
	}else{
		fmt.Println("修改失败")
	}
}
//查询
func sel(db *sql.DB ){
	stmt,err:=db.Query("select * from people")
if err!=nil{
	fmt.Println("预处理失败")
	return
}
	defer stmt .Close()
 for  stmt.Next(){
 	var id int
 	var address,name string
 	stmt .Scan(&id,&name,&address)//顺序不要搞错
 	fmt.Println(id,name,address)
 }  //这里的stmt直接变成了rows
}
