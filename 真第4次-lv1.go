package main

import (
	"fmt"
	"time"
)

func main() {
	var t []int64
	var a int64
	var i int
	fmt.Printf("请输入时间戳 （result结束）:")
	_,err:=fmt.Scanf("%v", &a)
	for err == nil {
		t = append(t, a)
		i++
		_,err= fmt.Scanf("%v", &a)
	}
	//获取输入
	dec(t, i)
}

//输出
func dec(ti []int64, n int) {
	for a := 0; a < n; a++ {
		fmt.Println("input OK!")
	}
	fmt.Println("the result are:")
	for a := 0; a < n; a++ {
		k := time.Unix(ti[a], 0)
		fmt.Println(k)
	}
}
