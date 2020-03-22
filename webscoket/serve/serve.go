package serve

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

//绑定一个send
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore { //如果这个和该用户的不同
			conn.Send <- message
		}
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		//准备好了
		case conn := <-manager.Register:
			//改为已经传过来
			manager.Clients[conn] = true
			jsonMessage, err := json.Marshal(&Message{Content: "/A new socket has connected."})
			if err != nil {
				log.Fatalf("json error:", err)
			}
			manager.send(jsonMessage, conn)
		//未准备好
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, err := json.Marshal(&Message{Content: "/A socket has disconnected."})
				if err != nil {
					log.Fatalf("json error:", err)
				}
				manager.send(jsonMessage, conn)

			}
			//传输开始，那边可以开协程
		case message := <-manager.Broadcast:

			for conn := range manager.Clients {
				select {

				case conn.Send <- message:{

				}

				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

//读取
func (c *Client) Read(manager ClientManager) {
	defer func() {
		manager.Unregister <- c
		c.Socket.Close()
	}()

	for {

		_, message, err := c.Socket.ReadMessage() //读取信息
		if err != nil {
			manager.Unregister <- c

			c.Socket.Close()
			break
		}
		//直接传过去
		jsonMessage, err := json.Marshal(&Message{Sender: c.Name , Content: string(message), Recipient: c.Recipient})
		if err != nil {
			log.Fatal("json error:", err)
		}

		manager.Broadcast <- jsonMessage
	}
}

//写入
func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {

				c.Socket.WriteMessage(websocket.CloseMessage, []byte("get message error"))
				return
			}
			fmt.Println(string(message))
			c.Socket.WriteMessage(websocket.TextMessage, message) //这里是输出message,以text
		}
	}
}
