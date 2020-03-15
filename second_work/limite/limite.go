package limite

import "log"

type ConnLimiter struct {
	concurrentConn int //个数
	bucket         chan int
}

func CreatChan()*ConnLimiter{
	return  &ConnLimiter{
		concurrentConn:100,
		bucket:make(chan int,100),
	}

}
//往管道里传
func (cl *ConnLimiter)GetConn()bool{
	//判断管道是否堵塞
	if len(cl.bucket)>=cl.concurrentConn{
		log.Println("Reached the rate limitation.")
		return  false
	}
	cl.bucket<-1
	return  true
}
//取出
func (cl *ConnLimiter)ReleaseCoon(){
	<-cl.bucket
	log.Println("release Conn")
}