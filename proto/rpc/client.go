package rpc

import (
	"log"
	"net"
	"reflect"
)

//客户端
type Client struct {
coon net.Conn
}

//创建函数
func NewClient(coon net.Conn)*Client{
	return  &Client{coon}
}

//实现通用的rpc客户端
//绑定方法  传入函数名

//具体实现再server端，client是函数原型
//使用mekefunc()完成原型到函数的调用

//fptr指向函数
//xxx.callRpc("",&)
func (c *Client)callRpc(rpcName string,fptr interface{}){
	//通过反射，获取fptr未初始化的函数模型
	fn:=reflect.ValueOf(fptr).Elem()
	//另一个函数，对第一个函数参数进行操作
	//完成Server的交互
	f:= func(args []reflect.Value) []reflect.Value{
		//处理
		inArgs:=make([]interface{},0,len(args))
		for _,v:=range args{
			inArgs=append(inArgs,v)
		}
		//创建链接
		clientSession:=NewSession(c.coon)
		//编码
		rpcRpc:=RPCDate{Name:rpcName,Args:inArgs}
		b, err := encode(rpcRpc)
		if err!=nil{
			log.Panicln(err)
		}
		//写
		err = clientSession.Write(b)
		if err!=nil{
			log.Panic(err)
		}
		//读
		r, err := clientSession.Read()
		if err!=nil{
			log.Panic(err)
		}
		//解码
		respdate, err := decode(r)
		if err!=nil{
			log.Println(err)
		}
		//处理服务端的返回结果
		outArgs:=make([]reflect.Value,0,len(respdate.Args ) )
		for i,arg:=range rpcRpc.Args{
			if arg==nil{
				//给一个真正的类型，不能是nil
				outArgs=append(outArgs,reflect.Zero(fn.Type().Out(i) )  )
				continue
			}
			outArgs=append(outArgs,reflect.ValueOf(arg)  )
		}
		return  outArgs
	}

	//参数1：未初始化的函数值，类型是reflect.Type
	//参数2：函数，对第一个函数参数操作
	//返回reflect.Value 类型
	//MakeFunc 使用传入的函数原型，创建一个绑定参数2的新函数
	v:=reflect.MakeFunc(fn.Type(),f)
	//赋值
	fn.Set(v)
}
