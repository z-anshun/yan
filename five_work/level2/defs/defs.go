package defs

type Information struct {
	Versio  string     `json:"versio"`   //"version": "2020.2.15",
	StuNum  string     `json:"stu_num"`  //"stuNum": "2018211219",
	NowWeek string     `json:"now_week"` //"nowWeek": 8,
	Success bool       `json:"success"`  //"success": true,
	Status  int        `json:"status"`   //"status": 200,
	Data    []Schedule `json:"data"`
}

type Schedule struct {
	Hash_day     int    `json:"hash_day"`     //"hash_day": 0,
	Hash_lesson  int    `json:"hash_lesson"`  //"hash_lesson": 0,
	Begin_lesson int    `json:"begin_lesson"` //"begin_lesson": 1,
	Course_num   string `json:"course_num"`   //"course_num": "A2031820",
	Day          string `json:"day"`          //"day": "星期一",
	Lesson       string `json:"lesson"`       //"lesson": "一二节",
	Course       string `json:"course"`       //"course": "网络营销",
	Teacher      string `json:"teacher"`      //"teacher": "徐秀珍",
	Classroom    string `json:"classroom"`    //"classroom": "2503",
	RawWeek      string `json:"raw_week"`     //"rawWeek": "1-16周",
	WeekModel    string `json:"week_model"`   //"weekModel": "all",
	WeekBegin    int    `json:"week_begin"`   //"weekBegin": 1,
	WeekEnd      int    `json:"week_end"`     //"weekEnd": 16,
	Type         string `json:"type"`         //"type": "必修",
	Period       string `json:"period"`       //"period": 2,
	Week         []int  `json:"week"`
}
