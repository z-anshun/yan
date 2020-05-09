package rpc

import (
	"log"
	"net"
	"reflect"
)

//声明服务端
type Server struct {
	addr string
	//服务端维护的函数名到函数的反射值
	funcs map[string]reflect.Value
}

//创建服务端对象
func NewServer(addr string)*Server{
	return  &Server{addr:addr,funcs: make(map[string]reflect.Value) }
}

//服务端绑定注册方法
func (s *Server)Register(rpcName string,f interface{} ){
	//如果有
	if _,ok:=s.funcs[rpcName];ok{
		return
	}
	//如果没有
	fVal:=reflect.ValueOf(f)
	s.funcs[rpcName]=fVal
}
//服务端等待调用
func (s *Server)Run(){
	//监听->链接->会话
	//监听
	lis,err:=net.Listen("tcp",s.addr)
	if err!=nil{
		log.Panicln("lister addr is err:",err)
		return
	}
	for{
		//拿到链接
		coon,err:=lis.Accept()
		if err!=nil{
			log.Println("accept err",err)
			return
		}
		//创建会话
		srvSession:=NewSession(coon)
		//读
		b,err:=srvSession.Read()
		if err!=nil{
			log.Println("read error:",err)
			return
		}
		//对数据解码
		rpcDate,err:=decode(b)
		if err!=nil{
			log.Println("decode err:",err)
			return
		}
		//根据读取到数据的name来调用函数名
		f,ok:=s.funcs[rpcDate.Name]
		if !ok{
			log.Println("this function not exist")
			return
		}
		//解析遍历客户端的参数
		inArgs:=make([]reflect.Value,0,len(rpcDate.Args))
		for _,v:=range rpcDate.Args {
			inArgs=append(inArgs,reflect.ValueOf(v))
		}
		//反射调用方法，传入参数
		out:=f.Call(inArgs)
		//解析遍历结构
		outArgs:=make([]interface{},0,len(out))
		for _,o:=range out {
			outArgs=append(outArgs,o)
		}
		//包装返回
		respRPCData:=RPCDate{rpcDate.Name,outArgs }
		//编码
		respBytes,err:=encode(respRPCData)
		if err!=nil{
			log.Println("encode err:",err)
			return
		}
		//写数据
		err=srvSession.Write(respBytes)
		if err!=nil{
			log.Println("session write error:",err)
			return
		}
	}
}
