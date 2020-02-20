package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

type apibody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

func main() {
	//get, _ := http.Get("http://127.0.0.1:8000/user")

	//res, _ := ioutil.ReadAll(get.Body)

	//m1 := make(map[string ]string ,1)
	//m1= map[string]string{
	//	"url":"http://www.127.0.0.1/use",
	//	"method":"123",
	//	"req_body":"22",
	//}
	//k,err:=json.Marshal(m1)
	//aip := &apibody{}
	//err=json.Unmarshal(k, aip)
	//fmt.Println(string(k),err,aip)
	k,_:=NewUUID()
	n:=k[0:2]=="Av"
	fmt.Println(n)

}
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return  fmt.Sprintf("Av%x-%x-%x-%x-%x",uuid[0:4],uuid[4:6],uuid[6:8],uuid[8:10],uuid[10:16]),err
}

