package first_work

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init(){
	//打开数据库
	open, err := gorm.Open("mysql", "root:123@(127.0.0.1:3306)/first?charset=utf8&parseTime=true")
	if err!=nil{
		log.Fatal("mysql open err:",err)
		return
	}
	DB=open
}