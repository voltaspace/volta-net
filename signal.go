package voltanet

import (
	"github.com/google/wire"
	"github.com/gorilla/websocket"
)
type Signal struct {

}

var SignalProvider = wire.NewSet(NewSignal)

func NewSignal() *Signal{
	return &Signal{}
}

/**
 * 当客户端连接时触发
 */
func (signal Signal) OnConnect(ws *websocket.Conn, clientId string) {

}

/**
 * 当客户端发来消息时触发
 */
func (signal Signal) OnMessage(ws *websocket.Conn, message string) {

}

/**
 * 当用户断开连接时触发
 */
func (signal Signal) OnClose(client ClientInfo) {

}