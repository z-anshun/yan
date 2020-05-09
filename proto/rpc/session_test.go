package rpc

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

func TestSession_Read(t *testing.T) {
	//定义监听ip和端口
	addr:="127.0.0.1:8000"
	my_data:="hello world"
	//等待组
	wg:=sync.WaitGroup{}
	wg.Add(2)
	//写
	go func() {
		defer  wg.Done()
		//创建tcp
		lis,err:=net.Listen("tcp",addr)
		if err!=nil{
			log.Panicln(err)
		}
		coon,_:=lis.Accept()
		s:=Session{coon:coon}
		err=s.Write([]byte(my_data))
		if err!=nil{
			log.Println(err)
		}

	}()
	//读
	go func() {
		defer  wg.Done()
		//创建tcp
		coon,err:=net.Dial("tcp",addr)
		if err!=nil{
			log.Panicln(err)
		}

		s:=Session{coon:coon}
		d,err:=s.Read()
		if err!=nil{
			log.Println(err)
		}
		fmt.Println(string(d) )
	}()
}
