package pachong

import (
	"encoding/json"
	"five_work/level2/defs"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func GetLesson() *[]defs.Schedule {

	str, code := getInformation("http://jwc.cqupt.edu.cn/kebiao/kb_rw.php")
	if code != 200 {
		return nil
	}

	all := strings.ReplaceAll(str, "\r\n", "")

	//理论  名字  老师  星期   教师
	les := `<td rowspan='1' align='center'>(.*?)</td>.*?<td rowspan='1'>(.*?)-(.*?)</td>.*?<td rowspan='1'>.*?</td>.*?<td rowspan='1'>选修</td><td rowspan='1' align='center'>全校选修<br> </td><td>(.*?)</td>.*?<td>(星期.*?)(第.*?节) (.*?周)</td><td>(.*?)</td>.*?<td rowspan='1' align='center'><a.*?>名单</a></td>`

	lesReg := regexp.MustCompile(les)
	l := lesReg.FindAllStringSubmatch(all, -1)
	//理论 A1115010 宇宙的奥秘 潘宇 星期2 第9-10节 1-8周 3205
	var classes []defs.Schedule
	for _, v1 := range l {
		var class defs.Schedule
		class.Type = "选修"
		class.Course_num = v1[2]
		class.Course = v1[3]
		class.WeekModel = "all"
		class.Teacher = v1[4]
		class.Day = v1[5]
		class.Lesson = v1[6]
		class.RawWeek = v1[7]
		class.Classroom = v1[8]
		weeks := getWeek(class.RawWeek)
		class.WeekBegin = weeks[0]
		class.WeekEnd = weeks[len(weeks)-1]
		date, err := json.Marshal(weeks)
		if err != nil {
			log.Println("json error:", err)
		}
		class.Period = "1"
		class.Week = string(date)
		classes = append(classes, class)

	}
	//regexp.MustCompile(``)
	return &classes
}
func getWeek(week string) []int {
	var weeks []int
	w := strings.Split(week, ",")
	if len(w) >= 2 {
		for _, v := range w {
			wi, err := strconv.Atoi(strings.Trim(v, "周"))
			if err != nil {
				log.Panic("string to int error")
				return nil
			}
			weeks = append(weeks, wi)
		}
	} else {
		w := strings.Split(week, "-")
		if len(w) < 2 {
			fmt.Println(w)
		}
		start, err := strconv.Atoi(w[0])
		if err != nil {
			log.Panic("string to int error")
			return nil
		}
		end, err := strconv.Atoi(strings.Trim(w[1], "周"))
		if err != nil {
			log.Panic("string to int error")
			return nil
		}
		for i := start; start <= end; i++ {
			weeks = append(weeks, i)
		}

	}
	return weeks
}
