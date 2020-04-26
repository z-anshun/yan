package db

import (
	"encoding/json"
	"five_work/level2/defs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123@(127.0.0.1:3306)/students?charset=utf8&parseTime=true")
	if err != nil {
		log.Panicln("open mysql error:", err)
	}
	//db.CreateTable(&defs.Information{})
	//db.CreateTable(&defs.Schedule{})
	DB = db
}

//添加学生到数据库
func AddStudents(s *defs.Information) {
	tx := DB.Begin()
	if err := tx.Create(s); err != nil {
		tx.Rollback() //错了就测回
	}
	tx.Commit()

}
func FindStudents(i interface{}) *defs.Information {
	var s defs.Information
	if len(i.(string)) != 0 {

		err := DB.Model(&defs.Information{}).Where("stu_num=?", i.(string)).Find(&s).Error
		if err != nil {
			log.Println("find students error:", err)
		}
	}
	return &s

}

//这里是增加选修了   接口
//知道num 和  星期几的
func AddLesson(i interface{}, l *defs.Schedule) {
	student := FindStudents(i)
	err := json.Unmarshal([]byte(student.Schedules ), &student.Data)
	if err != nil {
		log.Println("json error:", err)
		return
	}
	tx := DB.Begin()

	//dx_num_day_lesson_week  索引的顺序  星期几->第节课->哪几周  主键 course
	num := tx.Model(&defs.Schedule{}).Where(&defs.Schedule{Course: l.Course}).Select("course_num").Value
	if num == nil {
		tx.Rollback() //错了就测回
		return

	}
	if err := tx.Model(&defs.Schedule{}).Where("course_num=? and day=?", num, l.Day).Select("day", "lesson", "week").Find(&l).Error; err != nil {
		log.Println("get lesson eroor:", err)
		tx.Rollback() //错了就测回
		return
	}
	tx.Commit()
	for _, v := range student.Data {
		if v.Day == l.Day && v.Lesson == l.Lesson && !comWeek(v.Week, l.Week) {
			fmt.Println("选课重了！！！")
		}
	}
	//如果没重叠
	if err := DB.Model(&defs.Schedule{}).Where("course_num=? and day=?", num, l.Day).Find(&l).Error; err != nil {
		log.Println(err)
		return
	}
	str, _ := json.Marshal(append(student.Data, *l))
	if err := DB.Model(&student).Update("schedules", string(str)); err != nil {
		log.Println(err)
		return
	}

}
func comWeek(w1 string, w2 string) bool {
	var wk1, wk2 []int
	if err := json.Unmarshal([]byte(w1), &wk1); err != nil {
		log.Println(err)
		return false
	}
	if err := json.Unmarshal([]byte(w2), &wk2); err != nil {
		log.Println(err)
		return false
	}
	for _, v1 := range wk1 {
		for _, v2 := range wk2 {
			if v1 == v2 {
				return false
			}
		}
	}
	return true
}
func PutLesson(less *[]defs.Schedule) {
	//加入数据库
	for _, v := range *less {
		err := DB.Model(&defs.Schedule{}).Create(&v).Error
		if err != nil {
			fmt.Println("put lesson error:", err)
			return
		}
	}
}
