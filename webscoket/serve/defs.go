package serve

import "github.com/gorilla/websocket"

type Client struct {
	Name      string          `json:"name"`   //自己的名字2
	Socket    *websocket.Conn `json:"socket"` //这个是stock里面的东西
	Send      chan []byte     `json:"send"`
	Recipient string          `json:"recipient"` //接收者
}

type ClientManager struct {
	Clients    map[*Client]bool `json:"clients"`   //这个里面包含这发送的信息
	Broadcast  chan []byte      `json:"broadcast"` //存放的信息
	Register   chan *Client     `json:"register"`  //是否注册
	Unregister chan *Client     `json:"unregister"`
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

//初始化
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}
