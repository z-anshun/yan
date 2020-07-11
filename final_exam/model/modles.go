package model

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"

)

type User struct {

	Name     string `json:"name" gorm:"index:na_pa;type:varchar(256)"`
	Password string `json:"password" gorm:"index:na_pa;type:varchar(20)"`
	gorm.Model
}

//创建一个棋盘
type Chess struct {
	Board [15][15]int
}
type Message struct {
	Name   string          `json:"name"`
	Socket *websocket.Conn `json:"socket"` //这个是stock里面的东西
	Send   chan []byte     `json:"send"`
	RoomId int             `json:"room_id"` //房间号
	Round  bool            `json:"round"`   //为true才是自己的回合
	Black  bool            `json:"black"`   //自己是白棋还是黑棋
	//这里还有个有退出
	Exit bool `json:"exit"`
}
type UserMessage struct {
	Users      map[*Message]bool
	Register   chan *Message
	Unregister chan *Message //未登录
	Ch         Chess         //棋盘
	Other      chan int      //人数
	Defeat     chan *Message //胜利者
	//TimeKeep   time.Duration //计时，，时间戳,前端实现更好
	Record     string
}
type Inter struct {
	//0为交流 1认输  2求和  3悔棋
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Action  string `json:"action"`  //落子的地方 用,隔开 （x,y）
	Content string `json:"content"` //聊天
}

//创建一个通道
//一个用户->房间号->是否准备好

var UserMessages = make(map[int]*UserMessage)

func NewUsers() *UserMessage {
	return &UserMessage{
		Users:      make(map[*Message]bool),
		Register:   make(chan *Message),
		Unregister: make(chan *Message),
		Other:      make(chan int, 2),   //这个相当于人数
		Defeat:     make(chan *Message), //这个相当于人数
	}
}

//记录文件好久删除
type FileRecord struct {
	RoomId  int`json:"room_id" gorm:"primary_key;type:int(10)"`
	EndTime int`json:"end_time"`
	Record string`json:"record"gorm:"type:text"`
}
