package model

import "src/second_work/limite"

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var Final = false            //这是时间判断
var Coon *limite.ConnLimiter //管道
var T int64                  //开始的时间