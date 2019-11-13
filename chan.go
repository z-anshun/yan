package main

import (
	"fmt"
	"time"
)

var myres = make(map[int]int,20)
var ch = make(chan int, 20)

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second * 5)
	for _, v := range myres {
		fmt.Println(v)
	}
}
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch <- res //关锁
	myres[n] = <-ch
	//开锁
}
