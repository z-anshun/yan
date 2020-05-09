package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

//会话链接的结构体
type Session struct {
	coon net.Conn
}

func NewSession(coon net.Conn) *Session {
	return &Session{coon: coon}
}

//写数据
func (s *Session) Write(data []byte) error {
	buf := make([]byte, 4+len(data))
	//写入头部数据，记录长度
	//binary只认固定长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	//写入数据
	copy(buf[4:], data)
	_, err := s.coon.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
//从链接里读数据
func (s *Session)Read()([]byte,error){
	//读取头部
	header:=make([]byte,4)
	//readfull读取指定长度
	_, err := io.ReadFull(s.coon, header)
	if err!=nil{
		return  nil,err
	}
	//读取数据
	dateLen:=binary.BigEndian.Uint32(header)
	date:=make([]byte,dateLen)

	_,err=io.ReadFull(s.coon,date)
	if err!=nil{
		return  nil,err
	}
	return date,nil
}
