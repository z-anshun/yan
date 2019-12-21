package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"
)

/*
一个数据文本`Student.txt`,里面有按学号排序好的学生数据

你需要的是在这些学生数据中按照题目给定的信息（学号573结尾，名字两个字）找到指定的学生信息

### 加分项

- 多线程

### 题目要求

不能自己手动找出来*/
func main() {
	fliepath := "Students.txt"
	file, e := os.Open(fliepath)
	defer file.Close()
	if e != nil {
		fmt.Println("打开文件失败")
		return
	}
	reader := bufio.NewReader(file)
	for {
		s, e := reader.ReadString('}')
	if e ==io.EOF  {
		break
	}
	go find(s)
}
time.Sleep(time.Second*10)
}
func find(str string ){
	var count int //这个是拿来计算空格数，，也就是名字数
	k1:=`"xmEn":"(.*?)","xb".*?`//写两个正则匹配
	 k:=`[\s\S]*?"xh":"(\d*?)","xm":"(.*?)".*?`
	compile := regexp.MustCompile(k)
	compile1:= regexp.MustCompile(k1)

	s := compile.FindStringSubmatch(str)
	s1:=compile1.FindStringSubmatch(str)
	if len(s1)!=0 {
		for _, k := range s1[1] {
			if string(k) == " " {
				count++
			}
		}
	}
	if len(s)!=0 {
		if s[1][7:] == "573" && count == 2 {
			fmt.Println("找到了：")
			fmt.Println(s1[1])
		}
	}
}
