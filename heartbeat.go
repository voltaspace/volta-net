package voltanet

import (
	"github.com/google/wire"
	"github.com/gorilla/websocket"
	"time"
	"voltanet/constant/cmd"
	"voltanet/utils"
)


type WsHeartBeat struct {
	Events *Events
}

var WsHeartBeatProvider = wire.NewSet(NewWsHeartBeat)

func NewWsHeartBeat(events *Events) *WsHeartBeat{
	return &WsHeartBeat{events}
}

//.开启心跳检测
func (wsHeartBeat *WsHeartBeat) Run() {
	for {
		var outSocket map[string]ClientInfo = make(map[string]ClientInfo)
		//.扫描过期已绑定socket
		Client.Lock.RLock()
		for k, v := range Client.ClientList{
			var date int64 = utils.GetDate()
			if ((date - v.Pant) > cmd.PING_INTERVAL) {
				var client ClientInfo = *v
				outSocket[k] = client
			}
		}
		Client.Lock.RUnlock()
		//.清除已绑定socket
		Client.Lock.Lock()
		for k, v := range outSocket {
			v.WsConn.Close("[volta-network]:heart beat timeout close")
			delete(Client.ClientList, k) //.踢出socket队列
		}
		Client.Lock.Unlock()
		time.Sleep(time.Second * 3)
	}
}

func (wsHeartBeat *WsHeartBeat) HeartBeatUpdate(ws *websocket.Conn) bool {
	client,err := wsHeartBeat.Events.GetClientInfoBySocket(ws)
	if err != nil {
		return false
	}
	//.更新时间
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	if _,ok := Client.ClientList[client.Uid];ok{
		Client.ClientList[client.Uid].Pant = time.Now().Unix()
	}
	return true
}

