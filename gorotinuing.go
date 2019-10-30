package main

import "fmt"

var ch = make(chan int, 1)
var ch1 = make(chan bool, 1) //创建两个管道，一个读，一个写
func read(n int) {
	var bo bool = true
	for i := 2; i*i < n; i++ {
		if (n == 2) {
			break
		}
		if (n%i == 0) {
			bo = false
			break
		}
	} //判断是否为素数
	if (bo) {
		ch <- n //传入素数
	}
	if (n == 10000) {
		close(ch) //读到10000了，该关了
	}
}
func write() {
	for {
		a, ok := <-ch
		if !ok {
			ch1 <- true
			close(ch1)
			break
		}
		fmt.Println(a) //打印素数
	}
}
func main() {
	for i := 2; i <= 10000; i++ {
		go read(i)
		go write()
	} //开黑多个协程
	ch1 <- false
	for {
		_, ok := <-ch1
		if !ok {
			break
		}

	}

}
