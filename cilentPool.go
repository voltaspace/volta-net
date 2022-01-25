package voltanet

import (
	"sync"
)

type ClientBind struct {
	ClientList map[string]*ClientInfo
	Lock       sync.RWMutex
}

//.初始化客户端队列(map)
var clientMute sync.RWMutex

var Client = ClientBind{
	ClientList: make(map[string]*ClientInfo),
	Lock:       clientMute,
}

//.client结构
type ClientInfo struct {
	ClientId  string
	WsConn    *WsConn
	Pant      int64
	Session   string
	Uid       string
	GroupNm   string
}
