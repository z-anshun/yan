package main

import (
	"encoding/json"
	"five_work/level2/limite"
	"five_work/level2/pachong"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	c := limite.NewConn()
	for i := 2019210000; i < 2019251000; i++ {
		url := "http://jwc.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(i)

		c.PutConn()
		go Json(url, c)
		time.Sleep(time.Millisecond * 100)

	}

	fmt.Println("爬取完毕")

}

func Json(url string, coon *limite.ConnLimiter) {
	student := pachong.Deal(url)
	if student.Data != nil {
		s, err := json.Marshal(student)
		if err != nil {
			log.Println("json error:", err)
		}
		fmt.Println(string(s) )
		writInformation(s)


	}
	coon.ReleaseConn()

}
func writInformation(student []byte) {
	f, err := os.OpenFile("student.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		log.Println("OpenFile error:", err)
	}
	_, err = f.Write(student)
	if err != nil {
		log.Println("write schedule error:", err)
	}

}
