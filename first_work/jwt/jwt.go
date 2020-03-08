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

func NewHeader() Header {
	return Header{
		Api: "HS256",
		Typ: "JWT",
	}
}

func CreatJwt(name string, id int) string {
	h := NewHeader()
	p := PlayLoad{
		UserName: name,
		Id:       id,
		Iss:      "jwt",
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Exp:      strconv.FormatInt(time.Now().Add(time.Hour*20).Unix(), 10),
	}
	header, err := json.Marshal(&h) //转化为json
	if err != nil {
		log.Fatal("json err:", err)
	}
	playload, err := json.Marshal(&p)
	if err != nil {
		log.Fatal("json err:", err)
	}
	//转换
	headerbase64 := base64.StdEncoding.EncodeToString(header)
	playlodbase64 := base64.StdEncoding.EncodeToString(playload)
	str := strings.Join([]string{headerbase64, playlodbase64}, ".")

	key := "first"
	mac := hmac.New(sha256.New, []byte(key)) //创建一个hash
	mac.Write([]byte(str))                   //写入
	s := mac.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(s)
	//header+"."+playload+"."+signature
	token := str + "." + signature

	return token


}
func CheckToken(token string) (uid int, name string, err error) {
	arr := strings.Split(token, ".")
	if len(arr)!=3{
		return  0,"",errors.New("token error")
	}
	_, err = base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		return 0, "", err
	}
	p,err:=base64.StdEncoding.DecodeString(arr[1])
	if err!=nil{
		return  0,"",err
	}
	s,err:=base64.StdEncoding.DecodeString(arr[2])
	if err!=nil{
		return  0,"",err
	}
	str:=arr[0]+"."+arr[1]

	key := "first"
	mac := hmac.New(sha256.New, []byte(key)) //创建一个hash
	mac.Write([]byte(str))                   //写入
	s1 := mac.Sum(nil)

	if !hmac.Equal(s,s1){
		return  0,"",errors.New("check error")
	}

	var playload PlayLoad
	//json化
	if err=json.Unmarshal(p,&playload) ;err!=nil{
		log.Fatal("json error")
	}

	return playload.Id,playload.UserName,nil
}
