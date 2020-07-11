package db

import (
	"final_exam/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	//初始化 101.201.140.26
	open, err := gorm.Open("mysql", "root:123@(127.0.0.1:3306)/game?charset=utf8&parseTime=true")
	if err != nil {
		log.Println("open mysql error ",err)
	}
	e := open.HasTable(&model.User{})
	if !e {

		err := open.AutoMigrate(&model.User{}).Error
		if err != nil {
			log.Println("Creat Table error")
		}
	}
	e = open.HasTable(&model.FileRecord{})
	if !e {

		err := open.AutoMigrate(&model.FileRecord{}).Error
		if err != nil {
			log.Println("Creat Table error")
		}
	}

	DB = open
}
