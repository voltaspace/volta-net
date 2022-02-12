package voltanet

import (
	"github.com/google/wire"
	"github.com/gorilla/websocket"
	"github.com/voltaspace/volta-net/constant/cmd"
	"time"
)


type WsHeartBeat struct {
	Events *Events
}

var WsHeartBeatProvider = wire.NewSet(NewWsHeartBeat)

func NewWsHeartBeat(events *Events) *WsHeartBeat{
	return &WsHeartBeat{events}
}

//.heartbeat Run
func (wsHeartBeat *WsHeartBeat) Run() {
	go wsHeartBeat.Scan()
}

//.update volta-net options SetOptions
func (wsHeartBeat *WsHeartBeat) SetOptions(options *Options){
	return
}

//. scan timeout client Scan
func (WsHeartBeat *WsHeartBeat) Scan(){
	defer EndStack("heartbeat")
	for {
		var outSocket map[string]ClientInfo = make(map[string]ClientInfo)
		Client.Lock.RLock()
		for k, v := range Client.ClientList{
			var date int64 = time.Now().Unix()
			if ((date - v.Pant) > cmd.PING_INTERVAL) {
				var client ClientInfo = *v
				outSocket[k] = client
			}
		}
		Client.Lock.RUnlock()
		Client.Lock.Lock()
		for k, v := range outSocket {
			v.WsConn.Close("[volta-net]:heart beat timeout close")
			delete(Client.ClientList, k) //.踢出socket队列
		}
		Client.Lock.Unlock()
		time.Sleep(time.Second * 3)
	}
}

//. update active time HeartBeatUpdate
func (wsHeartBeat *WsHeartBeat) HeartBeatUpdate(ws *websocket.Conn) bool {
	client,err := wsHeartBeat.Events.GetClientInfoBySocket(ws)
	if err != nil {
		return false
	}
	client.Pant = time.Now().Unix()
	return true
}

