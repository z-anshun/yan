package jwt

//头部的
type Header struct {
	Api string `json:"api"`
	Typ string `json:"typ"`
}

//Playload
type PlayLoad struct {
	Iss      string `json:"iss"` //
	Exp      string `json:"exp"` //消失时间
	Iat      string `json:"iat"` //开始作用时间
	UserName string `json:"user_name"`

}
type Jwt struct {
	Header
	PlayLoad
	Signature string
}
