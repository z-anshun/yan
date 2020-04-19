package main

import (
	"encoding/json"
	"five_work/douban/defs"
	"five_work/douban/pachong"
	"fmt"
	"log"
	"strconv"
)

func main() {


	var movies []*[]defs.Movie
	for i := 0; i <= 250; i += 25 {
		str_i := strconv.Itoa(i)
		movie := pachong.Deal("https://movie.douban.com/top250?start=" + str_i + "&filter=")
		movies = append(movies, movie)
	}

	//fmt.Println(movie)
	for _, v := range movies {
		for _, v1 := range *v {
			////获取图片
			//get, err := http.Get(v1.Img)
			//if err != nil {
			//	fmt.Println("request img error")
			//}
			////读取图片
			//all, err := ioutil.ReadAll(get.Body)
			//if err != nil {
			//	fmt.Println("get img error")
			//}
			//v1.Img = string(all)
			marshal, err := json.Marshal(v1)
			if err != nil {
				log.Panicln("json error")
			}
			fmt.Println(string(marshal))
		}
	}

}
