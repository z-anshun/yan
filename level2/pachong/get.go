package pachong

import (
	"encoding/json"
	"five_work/level2/defs"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var User_Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:75.0) Gecko/20100101 Firefox/75.0"

//创建一个鉴证
var C = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{

			Timeout:   30 * time.Second, //限制建立TCP连接的时间
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second, // 限制 TLS握手的时间

		ResponseHeaderTimeout: 10 * time.Second, //限制读取response header的时间

		ExpectContinueTimeout: 1 * time.Second, //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。

	},
}

func getInformation(url string) (string, int) {

	//创建新的访问
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("request error:", err)
		return "", 0
	}
	//把user-agent加进去
	req.Header.Add("User-Agent", User_Agent)
	resp, err := C.Do(req)
	if err != nil {
		log.Println("get response error:", err)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("get schedule error:", err)
	}
	//关闭读取的body
	defer resp.Body.Close()

	return string(all), resp.StatusCode
}
func Deal(url string) *defs.Information {
	s, code := getInformation(url)
	t := time.Now().Format("2006-01-02")

	var student defs.Information
	//当前时间
	student.Versio = t
	if code != 200 {
		student.Success = false
		return nil
	}
	student.Success = true
	student.Status = code
	sch := strings.ReplaceAll(s, "\r\n", "")
	//学号 第几周
	I := `<li>〉〉.*?学年\d学期 学生课表>>(\d{10}).*?</li>[\s\S]*?今天是第 (\d) 周 星期`
	inf := match(I, sch)
	if len(inf) == 0 {
		log.Println("get information error:")
	}
	student.StuNum = inf[1]  //学号
	student.NowWeek = inf[2] //当前是第几周

	//分别获取1-10节的课
	//fmt.Println(student)

	classes := `<td style='font-weight:bold;'>\d、\d+节[\s\S]*?<tr style='text-align:center'>`

	clasReg := regexp.MustCompile(classes)
	cla := clasReg.FindAllString(sch, -1)
	//n-m节  地点 哪几周 老师 选修/必修
	class := `<div class='kbTd'.*?>.*?<br>.*?-(.*?)<br>地点：([\s\S]*?) [\s\S]*?<br>(.*?)<font.*?<span.*?>(.*?) (.*?) (.*?)学分`

	schedule := assign(cla, class)

	classEnd := `<td style='font-weight:bold;'>11、12节[\s\S]*?</table>`
	claEndReg := regexp.MustCompile(classEnd)
	claEnd := claEndReg.FindAllString(sch, -1)

	scheduleEnd := assign(claEnd, class)
	//判断最后一个是不是没得课
	if len(*scheduleEnd) != 0 {

		(*scheduleEnd)[0].Lesson = "第11-12节"
		(*scheduleEnd)[0].Begin_lesson = 10
		*schedule = append(*schedule, (*scheduleEnd)[0])
	}
	student.Data = *schedule
	return &student

}

//匹配
func match(reg string, str string) []string {
	strReg := regexp.MustCompile(reg)
	return strReg.FindStringSubmatch(str)
}

//给课表赋值
func assign(cla []string, class string) *[]defs.Schedule {

	var studentsClass []defs.Schedule
	for k, v := range cla {
		d := `<td(.*?)</td>`
		dReg := regexp.MustCompile(d)
		day := dReg.FindAllStringSubmatch(v, -1)
		for day, vd := range day {

			if vd[1] != " >" {

				r := regexp.MustCompile(class)
				all := r.FindAllStringSubmatch(vd[0], -1)
				if len(all) == 0 {
					continue
				}
				//进行赋值

				for _, v1 := range all {
					var studentClass defs.Schedule
					studentClass.Day = putDay(day)
					switch k {
					case 0:
						studentClass.Lesson = "第1-2节"
						studentClass.Begin_lesson = 1
					case 1:
						studentClass.Lesson = "第3-4节"
						studentClass.Begin_lesson = 3
					case 2:
						studentClass.Lesson = "第5-6节"
						studentClass.Begin_lesson = 5
					case 3:
						studentClass.Lesson = "第7-8节"
						studentClass.Begin_lesson = 7
					case 4:
						studentClass.Lesson = "第9-10节"
						studentClass.Begin_lesson = 9
					}

					studentClass.Course = v1[1]
					studentClass.Classroom = v1[2]
					studentClass.RawWeek = v1[3]
					studentClass.Teacher = v1[4]
					studentClass.Type = v1[5]
					studentClass.Period = v1[6]
					//看是否有三节连上那种
					thr := `3节连上`
					if check(thr, v1[0]) {

						switch k {
						case 0:
							studentClass.Lesson = "第1-3节"

						case 2:
							studentClass.Lesson = "第5-7节"

						case 4:
							studentClass.Lesson = "第9-11节"

						}
					}
					//4节连上的。。恐怖
					four := `4节连上`
					if check(four, v1[0]) {

						switch k {
						case 0:
							studentClass.Lesson = "第1-4节"

						case 2:
							studentClass.Lesson = "第5-8节"

						case 4:
							studentClass.Lesson = "第9-12节"

						}
					}

					//提取数字
					i := `(\d+)`
					iReg := regexp.MustCompile(i)
					week := iReg.FindAllStringSubmatch(v1[3], -1)
					studentClass.WeekBegin, _ = strconv.Atoi(week[0][0])

					studentClass.WeekEnd, _ = strconv.Atoi(week[len(week)-1][0])

					//赋值week
					odd := `单`
					even := `双`
					con := `-`
					flag_odd := match(odd, v1[3])
					flag_even := match(even, v1[3])
					flag_con := match(con, v1[3])
					//flag 0  单,双周1
					flag := 0
					var allweek []int
					//判断单双周
					if len(flag_odd) != 0 || len(flag_even) != 0 {
						flag = 1
					}
					//判断是不是连续上
					if len(flag_con) != 0 {
						s, _ := strconv.Atoi(week[0][0])
						e, _ := strconv.Atoi(week[1][0])

						allweek = putWeek(s, e, flag)
						//判断后面还有不有
						if len(week) > 2 {
							for k, w := range week {
								if k >= 2 {
									w_i, _ := strconv.Atoi(w[0])
									allweek = append(allweek, w_i)
								}
							}
						}
					} else {
						for _, w := range week {
							w_i, _ := strconv.Atoi(w[0])
							allweek = append(allweek, w_i)
						}
					}
					w, err := json.Marshal(allweek)
					if err != nil {
						log.Println("json week error", err)
					}
					studentClass.Week = string(w)
					studentsClass = append(studentsClass, studentClass)

				}
			}
		}

	}

	return &studentsClass
}
func putDay(k int) string {
	switch k {
	case 1:
		return "星期1"
	case 2:
		return "星期2"
	case 3:
		return "星期3"
	case 4:
		return "星期4"
	case 5:
		return "星期5"
	default:
		return ""
	}

}

//加入哪几周的函数
func putWeek(start int, end int, flag int) []int {
	var week []int

	switch flag {
	case 0:
		{
			for ; start <= end; start++ {
				week = append(week, start)
			}
		}
	case 1:
		for ; start <= end; start += 2 {
			week = append(week, start)
		}
	}
	return week
}

func check(reg string, str string) bool {
	r := regexp.MustCompile(reg)
	if len(r.FindAllStringSubmatch(str, -1)) != 0 {
		return true
	}
	return false
}
