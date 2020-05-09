package rpc

import (
	"bytes"
	"encoding/gob"
)

//定义数据格式和编解码
type RPCDate struct {
	//函数名
	Name string
	//参数
	Args []interface{}
}
//定义交互格式

//编码
func encode(data RPCDate)([]byte,error){
	var buf bytes.Buffer
	//得到字节数组的编码器
	bufEnc:=gob.NewEncoder(&buf)
	//对数据编码
	if err:=bufEnc.Encode(data);err!=nil{
		return  nil,err
	}
	return  buf.Bytes() ,nil
}
//解码
func decode(b []byte)(RPCDate ,error){
	buf:=bytes.NewBuffer(b)
	//返回字节解码数组
	bufDec:=gob.NewDecoder(buf)
	var data RPCDate
	if err:=bufDec.Decode(&data);err!=nil{
		return data ,err
	}
	return  data,nil
}