package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Student struct {
	Number int    //`json:"number"`
	Name   string // `json:"name"`
	Times  int    //`json:"times"`
}

var students []Student //学生太多，多创几个切片
//var ch=make(chan string,1)
var (
	reName = `<li>〉〉2019-2020学年1学期 学生课表>>20\d{8}\S{1,}`
)

func main() {
	filep := "D:/name.txt"
	for m := 2016000000; m < 2020000000; m++ {
		k := strconv.Itoa(m)                                       //学号转换为字符串
		ht := "http://jwzx.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + k //得到每个学生的网址

		//time.Sleep(time.Second*3) //让文件写好

		go Spy(ht) //开启协程,并传递地址清空文件

		time.Sleep(time.Microsecond * 100) //，防止出事

	} //读取未读完的
	time.Sleep(time.Second * 60)
	file, _ := os.Open(filep) //打开文件
	reader := bufio.NewReader(file)
	for { //读取
		s, ok := reader.ReadString('}') //一行一行的读取
		//fmt.Printf("%s",s)
		var sil Student
		json.Unmarshal([]byte(s), &sil) //反序列化
		//if err != nil {
		//fmt.Println("defeat~~", err)
		//return
		//	}
		students = append(students, sil) //放入切片
		if ok == io.EOF {
			break
		}
	}
	//开始找重名的人
	Find()
	//比较并打印
	Compar()

}
func Spy(ht string) {
	resp, err := http.Get(ht) //爬取网站
	if err != nil {
		//fmt.Println("false err=", err)
		return
	} //判断
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error:talking is cheap", resp.StatusCode)
		return
	} //判断

	details, err := ioutil.ReadAll(resp.Body) //拿出dody
	html := string(details)
	if err != nil {
		fmt.Println(err)
	} //判断

	re := regexp.MustCompile(reName)          //正则匹配
	str := re.FindAllStringSubmatch(html, -1) //拿取所有
	if str == nil {
		return
	}
	for i, x := range str {

		if i >= 1 && x[i] == x[i-1] {
			continue //若重复则再选
		}
		for _, a := range x {
			if a != "" {//除去空的
				//调用函数，转到文件里
				Writ(a)
			}
		}

	}
}

//写入文件
func Writ(str string) {

	name := fmt.Sprintln(str [57:])
	name = strings.Trim(name, "\n  </li>") //剔除不需要的，留下名字
	if name==" "{
		return
	}
	number_str := fmt.Sprintf(str [47:57])
	number, _ := strconv.Atoi(number_str)
	var studen Student
	{
		studen.Name = name
		studen.Number = number
		studen.Times = 1
	}

	//students=append(students, studen)
	//fmt.Println(s)
	filepath := "D:/name.txt" //给出文件路径

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND, 0666) //打开

	if err != nil {
		fmt.Println("false  ", err)
		return
	} //判断
	defer file.Close() //关闭文件
	Stu, err := json.Marshal(&studen)
	if err != nil {
		return
	}
	write := bufio.NewWriter(file)
	write.WriteString(string(Stu))
	write.Flush()
}
func Find() {
	for i := 0; i < len(students); i++ {
		for k := i + 1; k < len(students); k++ { //逐个比较，并将次数加一

			if students[i].Name == students[k].Name {
				students[i].Times = students[i].Times + 1

			}
		}
	}
}
func Compar() {
	var stu Student
	for i := 0; i < len(students); i++ {
		for k := i + 1; k < len(students); k++ { //逐个比较打印出次数最多那个

			if students[i].Times > students[k].Times {
				stu = students[i]
			} else if students[i].Times < students[k].Times {
				stu = students[k]
			} else if students[i].Times == students[k].Times {
				continue

			}
		}
	}
	fmt.Printf("重名最多的同学是%s   一共有%d次\n", stu.Name, stu.Times)

}
