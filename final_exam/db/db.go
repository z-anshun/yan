package db

import (
	"database/sql"
	"final_exam/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	//初始化
	open, err := gorm.Open("mysql", "root:123@(localhost:3306)/game?charset=utf8&parseTime=true")
	if err != nil {
		log.Panicln("open mysql error")
	}
	_, err = open.Where(&model.User{}).Rows()
	if err != nil || err == sql.ErrNoRows {

		err := open.AutoMigrate(&model.User{}).Error
		if err != nil {
			log.Println("Creat Table error")
		}
	}
	_, err = open.Where(&model.FileRecord{}).Rows()
	if err != nil || err == sql.ErrNoRows {

		err := open.AutoMigrate(&model.FileRecord{}).Error
		if err != nil {
			log.Println("Creat Table error")
		}
	}

	DB = open
}
