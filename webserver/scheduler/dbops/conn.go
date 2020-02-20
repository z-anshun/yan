package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql",
		"root:123@tcp(localhost:3306)/bli?charset=utf8") //这里不会真正连接
	if err != nil {
		panic(err.Error())
	}
}
