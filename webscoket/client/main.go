package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	go timeWriter(conn)
	//defer  conn.Close()
	//循环读取，读得耍
	for {
		_, message, _ := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)

		}
		if len(message) != 0 {
			fmt.Printf("received: %s\n", message)
		}
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		//每3分钟报一下时间
		time.Sleep(time.Second * 60 * 3)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}
