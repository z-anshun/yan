package main

import "log"

//流控 chan->request->response->chan bucket token算法
type ConnLimiter struct {
	concurrentConn int //个数
	bucket         chan int
}

//构造函数
func NewConnLimiter(cc int)*ConnLimiter{
	return  &ConnLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),
	}
}

func(cl *ConnLimiter )GetConn()bool{
	if len(cl.bucket)>=cl.concurrentConn{
		log.Println("Reached the rate limitation.")
		return  false
	}
	cl.bucket<-1
	return  true
}
func(cl *ConnLimiter )ReleaseConn(){
	c:=<-cl.bucket
	log.Println("New connction coming:",c)
}