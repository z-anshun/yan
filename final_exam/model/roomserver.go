package model

import (
	"encoding/json"
	"final_exam/judge"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"strings"
)

func (m *UserMessage) send(message []byte, ignore *Message) {
	flag := false
	for coon := range m.Users {

		//这里是把信息发给对手
		if coon != ignore {
			//找到
			flag = true
			coon.Send <- message

		}
	}

	if !flag {
		jsonMsg, _ := json.Marshal(&Inter{Content: "您还没有对手"})
		//fmt.Println(ignore)
		ignore.Send <- jsonMsg
		fmt.Println("set=",string(message))
	}
}
func (m *UserMessage) Start() {
	for {
		select {
		case coon := <-m.Register:
			//未准备好
			if !m.Users[coon] {
				msg, _ := json.Marshal(&Inter{Content: "您的对手还未准备好", Code: 0})
				m.send(msg, coon)
			}
			//准备好了
			msg, _ := json.Marshal(&Inter{Content: "您的对手准备好了", Code: 0})
			m.send(msg, coon)
			//这里就是有人退出
		case coon := <-m.Unregister:
			{
				<-m.Other
				coon.Exit = true
				m.Users[coon] = false
				//这里有个决策
				msg, _ := json.Marshal(&Inter{Content: "您的对手退出了房间了", Code: 0})
				m.send(msg, coon)
			}
		case coon := <-m.Defeat:
			msg, _ := json.Marshal(&Inter{Content: "失败", Code: -1})
			coon.Send <- msg
			jsonMsg, _ := json.Marshal(&Inter{Content: "胜利", Code: -1})
			m.send(jsonMsg, coon)
		}
	}

}

//读取
//这里既要传给自己，也要传给别个
//这里下棋的判断不一样 0 棋盘  1 黑子  2白子
func (m *Message) Read(users *UserMessage) {
	defer func() {
		users.Unregister <- m
		m.Socket.Close()
	}()

	for {

		//假设那边传过来的也是一个json的
		_, p, err := m.Socket.ReadMessage()
		//fmt.Println(err)
		if err != nil {
			//退出了
			//等10秒，看是否是断线
			//这里应该还有个判断的
			//time.Sleep(time.Second * 10)

				fmt.Println("gg")
				users.Unregister <- m

				break


		}

		//不管读了什么，，塞进去
		users.Record += string(p)

		var msg Inter
		err = json.Unmarshal(p, &msg)
		//fmt.Println("msg=",msg)
		if err != nil {
			fmt.Println(string(p))
			marshal, _ := json.Marshal(&Inter{Content: "指令错误", Name: m.Name})
			m.Send <- marshal
		}
		//自己的
		//fmt.Println(msg.Name,m.Name)
		if msg.Name == m.Name {
			switch msg.Code {
			case 0:
				//自己下棋之前，对方要准备
				for _, v := range users.Users {
					if !v {
						jsonMsg, _ := json.Marshal(&Inter{Code: 1, Content: "有人未准备"})
						m.Send <- jsonMsg
					}
				}
				//未到自己回合就下棋
				if len(msg.Action) != 0 && !m.Round {
					m.Socket.WriteMessage(websocket.TextMessage, []byte("未到你的回合"))
				} else if len(msg.Action) != 0 {

					a := strings.Split(msg.Action, ",")
					x, err := strconv.Atoi(a[1])
					if err != nil {
						fmt.Println(err)
					}
					y, err := strconv.Atoi(a[1])
					if err != nil {
						fmt.Println(err)
					}
					//这是自己的步
					if m.Black {
						users.Ch.Board[x][y] = 1
						//判断是否禁手
						str := judge.Forbid(&users.Ch.Board, x, y)
						if len(str) != 0 {
							jsonMsg, _ := json.Marshal(&Inter{Content: str, Code: 0})
							m.Send <- jsonMsg
							users.Defeat <- m
						}
					} else {
						users.Ch.Board[x][y] = -1
					}
					m.Round = false

				}
				str, _ := json.Marshal(&Inter{Name: msg.Name, Content: msg.Content})
				//自己的话
				m.Send <- str
			case 1:
				jsonMsg, _ := json.Marshal(&Inter{Code: 1, Content: "您请求认输"})
				m.Send <- jsonMsg
			case 2:
				jsonMsg, _ := json.Marshal(&Inter{Code: 2, Content: "您请求求和"})
				m.Send <- jsonMsg
			case 3:
				jsonMsg, _ := json.Marshal(&Inter{Code: 3, Content: "您希望悔棋"})
				m.Send <- jsonMsg
			}
		}
		//正常转换
		//对方
		if m.Name != msg.Name {
			m.Send<-p
			switch msg.Code {
			case 0: //对方的
				users.send(p, m)
			case 1:
				jsonMsg, _ := json.Marshal(&Inter{Code: 1, Content: "对方请求认输"})
				users.send(jsonMsg, m)
			case 2:
				jsonMsg, _ := json.Marshal(&Inter{Code: 2, Content: "对方请求求和"})
				users.send(jsonMsg, m)
			case 3:
				jsonMsg, _ := json.Marshal(&Inter{Code: 3, Content: "对方希望悔棋"})
				users.send(jsonMsg, m)
			}
		}

	}
}

//写和做
func (m *Message) WriteAndDo(users *UserMessage) {
	defer func() {
		m.Socket.Close()
	}()
		//fmt.Println("write")
	for {
		select {
		case message, ok := <-m.Send:
				fmt.Println("send=",string(message))
			if !ok {

				m.Socket.WriteMessage(websocket.CloseMessage, []byte("get message error"))
				return
			}
			var action Inter
			//读取的时候处理了，，现在就不用处理了
			json.Unmarshal(message, &action)

			if !m.Round && len(action.Action) != 0 {
				//不在自己的回合，才能做 这是对方的
				a := strings.Split(action.Action, ",")
				x, err := strconv.Atoi(a[1])
				if err != nil {
					fmt.Println(err)
				}
				y, err := strconv.Atoi(a[1])
				if err != nil {
					fmt.Println(err)
				}
				//因为是对方发过来
				if !m.Black {
					users.Ch.Board[x][y] = 1
				} else {
					users.Ch.Board[x][y] = -1
				}
				m.Round = true

			}
			m.Socket.WriteMessage(websocket.TextMessage, message) //这里是输出message,以text
			//code为-1的时候就代表比赛结束
			if action.Code == -1 {
				break
			}
		}
	}

}
