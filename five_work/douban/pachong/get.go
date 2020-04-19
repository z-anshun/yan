package pachong

import (
	"five_work/douban/defs"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var user_Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:75.0) Gecko/20100101 Firefox/75.0"

func get(url string) []byte {
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("new request error:", err)
	}
	req.Header.Add("User-Agent", user_Agent)
	resp, err := c.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Println("http request error:", err)
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("get body error:", err)
	}
	return all
}
func Deal(url string) *[]defs.Movie {
	all := get(url)
	//先替换

	str := strings.ReplaceAll(string(all), "\n", "")
	//正则转化
	olReg := regexp.MustCompile(`<ol class="grid_view(.*?)</ol>`)
	//查找 ol
	ol := olReg.FindAllStringSubmatch(str, -1)
	liReg := regexp.MustCompile(`<li>(.*?)</li>`)
	li := liReg.FindAllString(ol[0][0], -1)

	//名称 图片网址  网址 导演 评分 评语
	imgReg := regexp.MustCompile(`<img width="\d+" alt="(.*?)" src="(.*?)"[\d\D]*?<div class="hd">[\s\S]*?<a href="(.*?)"[\d\D]*?导演: (.*?)\b[\d\D]*?<span class="rating_num" property="v:average">(.*?)</span>[\s\S]*?<span[ class="inq">(.*?)</span>]`)
	m := []defs.Movie{}
	for _, v := range li {
		s := imgReg.FindStringSubmatch(v)
		if len(s)==0{
			fmt.Println(v)
			fmt.Println("get information error")
		}
		var mi defs.Movie
		//不一定有评语
		if len(s)==6 {
			mi = defs.Movie{s[3], s[1], s[2], s[4], s[5], ""}
		}else{
			mi = defs.Movie{s[3], s[1], s[2], s[4], s[5], s[6]}
		}
		m = append(m, mi)
	}
		return  &m
}
