package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

func(j *Jwt)CreatJwt(name string,id string )string{
	j.PlayLoad=PlayLoad{
		Iss:"jwt",
		UserName:name,
		Iat:strconv.FormatInt(time.Now().Unix(),10),
		Exp:strconv.FormatInt(time.Now().Add(time.Hour*12).Unix() ,10) ,

	}
	j.Header=Header{
		Api: "HS256",
		Typ: "JWT",
	}

	h,err:=json.Marshal(&j.Header)
	if err!=nil{
	log.Fatal("json heard err:",err)
	}
	p,err:=json.Marshal(&j.PlayLoad)
	if err!=nil{
	log.Fatal("json playload err:",err)
	}
	//base64转换
	hbase64 := base64.StdEncoding.EncodeToString(h)
	pbase64:=base64.StdEncoding.EncodeToString(p)

	str:=strings.Join([]string{hbase64,pbase64},".")

	key:="jwt"
	mac:=hmac.New(sha256.New,[]byte(key))
	mac.Write([]byte(str))                   //写入
	s := mac.Sum(nil)

	j.Signature=base64.StdEncoding.EncodeToString(s)//转换

	token:=str+"."+j.Signature
	return token

}

func(j *Jwt)CheckJwt(token string)error{
	arr:=strings.Split(token,".")
	if len(arr)!=3{
	return  errors.New("token member error")
	}
	p,err:=base64.StdEncoding.DecodeString(arr[1])
	if err!=nil{
		return  errors.New("token error")
	}
	s,err:=base64.StdEncoding.DecodeString(arr[2])
	if err!=nil{
		return  errors.New("token error")
	}

	str:=strings.Join([]string{arr[0],arr[1]},".")

	key:="jwt"
	mac:=hmac.New(sha256.New,[]byte(key))
	mac.Write([]byte(str))                   //写入
	s1 := mac.Sum(nil)

	if !hmac.Equal(s,s1){
		return  errors.New("check error")
	}

	if err = json.Unmarshal(p, &j.PlayLoad);err!=nil{
		return  errors.New("check error")
	}

	return  nil

}
