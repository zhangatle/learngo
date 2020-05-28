package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
)

type ServerController struct {
	beego.Controller
}

func (c *ServerController) Get() {
	name := c.GetString("name")
	if len(name) == 0 {
		logs.Alert("name is null")
		c.Redirect("/", 302)
		return
	}
	logs.Info("get name:" + name + ", and send to chatRoom.html")
	c.Data["name"] = name
	c.TplName = "chatRoom.html"
}

func (c *ServerController) WS() {
	name := c.GetString("name")
	if len(name) == 0 {
		logs.Error("name is null")
		c.Redirect("/", 302)
		return
	}
	conn, err := (&websocket.Upgrader{}).Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		logs.Error("not a websocket connection")
		http.Error(c.Ctx.ResponseWriter, "not a websocket handshake", 400)
		return
	} else if err != nil {
		logs.Error("cannot setup websocket connection:", err)
		return
	}
	var client Client
	client.name = name
	client.conn = conn

	// 如果用户列表中没有该用户
	if !clients[client] {
		join <- client
		logs.Info("user:", client.name, "websocket connect success!")
	}

	defer func() {
		leave <- client
		err = client.conn.Close()
		if err != nil {
			logs.Info("disconnect error")
		}
	}()

	// 由于websocket一旦连接，便可以保持长时间通讯，则该接口函数可以一直运行下去，直到连接断开
	for {
		_, msgStr, err := client.conn.ReadMessage()
		if err != nil {
			break
		}
		logs.Info("Ws ------receive:" + string(msgStr))
		var msg Message
		msg.Name = client.name
		msg.EventType = 0
		msg.Message = string(msgStr)
		message <- msg
	}
}
