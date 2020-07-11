package db

import (
	"errors"
	"final_exam/model"
)
//登录
func GetUser(user *model.User) error {
	var u model.User
	DB.Model(&model.User{}).Where("name=?",user.Name).Scan(&u)

	if u.Password==user.Password{
		return nil
	}

	return errors.New("password error")
}
//注册
func AddUser(user *model.User)error{
	err:=DB.Where(&model.User{}).Create(user).Error
	if err!=nil{
		return err
	}
	return nil
}
//查重
func FindUser(n string)error{
	var u model.User
	err:=DB.Model(&model.User{}).Where("name=?",n).Scan(&u).Error
	if err!=nil||len(u.Name)==0{
		return nil
	}
	return errors.New("The name is repeat")
}
//添加记录
func AddRecord(r model.FileRecord)error{
	err:=DB.Model(&model.FileRecord{}).Create(r).Error
	if err!=nil{
		return err
	}
	return nil
}
//删除文件
func Del(id int)error{
	err:=DB.Model(&model.FileRecord{}).Delete("room_id=?",id).Error
	if err!=nil{
		return err
	}
	return nil
}