package db

import (
	"database/sql"
	. "five_work/douban"
	"five_work/douban/defs"
	"fmt"
	"log"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/douban")
	if err != nil {
		log.Println("open mysql error")
		return
	}
	DB = db

}

type M defs.Movie
//存入数据库
func (movie *M) Add() {
	DB.Ping()
	stmt, err := DB.Prepare("insert into movie values(?,?,?,?,?)")
	if err != nil {
		log.Println("prepare error:", err)
	}
	r, err := stmt.Exec(movie.Name, movie.Img, movie.Director, movie.Evaluation, movie.Comments)
	if err!=nil{
		log.Println("insert error:",err)
	}
	count, err := r.RowsAffected()
	if err != nil {
		fmt.Println("结果获取失败", err.Error())
	}
	if count > 0 {
		fmt.Println("新增成功！")
	} else {
		fmt.Println("新增失败!")
	}
	defer DB.Close()
}
func (movid *M)P(){
	fmt.Printf("test")
}