package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	name string
}

type Message struct {
	EventType byte   `json:"type"`
	Name      string `json:"name"`
	Message   string `json:"message"`
}

var (
	// 此处要设置缓冲通道，因为这是goroutine自己从通道中发送并接收数据
	// 若是无缓冲的通道，该goroutine发送数据到通道后被锁定，需要数据被接受后才能解锁，而接受数据的恰恰是它自己
	join    = make(chan Client, 10)  //用户接入通道
	leave   = make(chan Client, 10)  //用户退出通道
	message = make(chan Message, 10) //消息通道
	clients = make(map[Client]bool)  //用户映射
)

func init() {
	go broadcaster()
}

// 广播消息
func broadcaster() {
	for {
		// 哪个case可以执行，则转入到该case,都不可执行，则堵塞
		select {
		case msg := <-message:
			str := fmt.Sprintf("broadcaster------%s send message:%s\n", msg.Name, msg.Message)
			logs.Info(str)
			for client := range clients {
				data, err := json.Marshal(msg)
				if err != nil {
					logs.Error("Fail to marshal message:", err)
					return
				}
				if client.conn.WriteMessage(websocket.TextMessage, data) != nil {
					logs.Error("Fail to write message")
				}
			}
		case client := <-join:
			str := fmt.Sprintf("broadcaster---- %s join in the chat room\n", client.name)
			logs.Info(str)
			clients[client] = true
			var msg Message
			msg.Name = client.name
			msg.EventType = 1
			msg.Message = fmt.Sprintf("%s join in, there are %d persion in room", client.name, len(clients))
			message <- msg
		case client := <-leave:
			str := fmt.Sprintf("broadcaster-----------%s leave the chat room\n", client.name)
			logs.Info(str)
			// 如果该用户已经被删除
			if !clients[client] {
				logs.Info("the client had leaved, client's name:" + client.name)
				break
			}

			delete(clients, client) // 将用户从映射中删除

			// 将用户退出消息放入消息通道
			var msg Message
			msg.Name = client.name
			msg.EventType = 2
			msg.Message = fmt.Sprintf("%s leave, there are %d preson in room", client.name, len(clients))
			message <- msg
		}
	}
}
