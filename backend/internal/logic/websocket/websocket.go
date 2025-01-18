package logs

import (
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

type sWebsocket struct {
}

type WsMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	service.RegisterWebsocket(&sWebsocket{})
}

func New() service.IWebsocket {
	return &sWebsocket{}
}

func (s *sWebsocket) Upgrade(r *ghttp.Request) error {
	conn, err := upgrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		return err
	}
	for {
		var recv WsMessage
		err := conn.ReadJSON(recv)
		if err != nil {
			return err
		}

		var send WsMessage
		send.Type = "echo"
		send.Data = recv.Data
		err = conn.WriteJSON(send)
		if err != nil {
			return err
		}
	}
}
