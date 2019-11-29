package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var res = `[\s\S]*?学生课表>>(\d*?)(\D*?)  </li>[\s\S]*?`
//开内网后，又不能用外网，所以先爬进文件，再进行选录入数据库
func main() {
	for i := 2019211649; i <= 2019211688; i++ {
		k := strconv.Itoa(i)
		ht := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + k
		spy(ht)
	}
}
func spy(s string) {
	r, _ := http.Get(s)               //爬网站
	d, _ := ioutil.ReadAll(r.Body) //读取所有的内容
	detail := string(d) //转为string
	com := regexp.MustCompile(res)
	str := com.FindStringSubmatch(detail)
	file := "students"
	f, _ := os.OpenFile(file, os.O_APPEND|os.O_CREATE , 3360) //打开文件
	n, err := f.WriteString(str[1]+str[2]+"\n")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if n!=0{
		fmt.Println("ok")
	}
	defer f.Close()
}
