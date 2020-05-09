package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"seven_work/proto"
)

func main(){
	//监听
	lis,err:=net.Listen("tcp",":8081")
	if err!=nil{
		log.Panic("error:",err)
	}
	//创建服务端
	rpcServer:=grpc.NewServer()

	//创建实现方法
	proto.RegisterRegisterServer(rpcServer,new(proto.UserService))

	go func() {
		if err:=rpcServer.Serve(lis);err!=nil{
			fmt.Println(err.Error())
		}
	}()

	select {}

}
