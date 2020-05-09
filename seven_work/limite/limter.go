package limite

import "log"


type ConnLimiter struct {

	bucket         chan int
}

//构造函数
func NewConnLimiter()*ConnLimiter{
	return  &ConnLimiter{

		bucket:make(chan int,1),
	}
}

func(cl *ConnLimiter )GetConn(){

	cl.bucket<-1
	return
}
func(cl *ConnLimiter )ReleaseConn(){
	c:=<-cl.bucket
	log.Println("New connction coming:",c)
}