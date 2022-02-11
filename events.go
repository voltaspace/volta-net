package voltanet

import (
	"encoding/json"
	"errors"
	"github.com/google/wire"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/voltaspace/volta-net/wspl"
	"time"
)

type Events struct {
}

var EventsProvider = wire.NewSet(NewEvents)

func NewEvents() *Events {
	return &Events{}
}

//.查询客户端信息
func (wsEvents Events) GetClientInfo(clientId string) (res ClientInfo, err error) {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	if _, ok := Client.ClientList[clientId]; !ok {
		return res, errors.New("no key")
	}
	res = *Client.ClientList[clientId]
	return res, nil
}

//.通过uid获取绑定信息
func (wsEvents Events) GetClientInfoByUid(uid string) (res ClientInfo, err error) {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	if _, ok := Client.ClientList[uid]; !ok {
		return res, errors.New("no key")
	}
	res = *Client.ClientList[uid]
	return
}

//.获取所有客户端
func (wsEvents Events) GetAllClient() (clients []ClientInfo) {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	for _, v := range Client.ClientList {
		var client ClientInfo
		client = *v
		clients = append(clients, client)
	}
	return clients

}

//.同步uidlist获取对应通道
func (wsEvents Events) GetClientInfoByUids(uidlist []string) (clientList []ClientInfo, err error) {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	for _, v := range uidlist {
		if _, ok := Client.ClientList[v]; ok {
			clientList = append(clientList, *Client.ClientList[v])
		}
	}
	return
}

//.通过ws指针查找client信息
func (wxEvnets Events) GetClientInfoBySocket(ws *websocket.Conn) (clinet ClientInfo, err error) {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	for _, v := range Client.ClientList {
		if v.WsConn.Conn == ws {
			clinet = *v
			return
		}
	}
	return clinet, errors.New("GetClientInfoBySocket not found *websocket.Conn")
}

//.强制断开uid
func (wsEvents Events) CloseSocketByUid(uid string) (err error) {
	clientInfo, err := wsEvents.GetClientInfoByUid(uid)
	if err != nil {
		return
	}
	clientInfo.WsConn.Close("driving close")
	return
}

//.client绑定uid
func (wsEvents Events) Bind(session string, uid string, wsConn *WsConn) (clientId string, err error) {
	time.Sleep(time.Millisecond * 500) //.防止android握手未完成
	if session == "" || uid == "" {
		return "", errors.New("[volta-gateway]:bind client sesion&uid nil")
	}
	goId := uuid.NewV4()
	clientId = goId.String()
	var date int64 = time.Now().Unix()
	var bindInfo = &ClientInfo{
		ClientId: clientId,
		WsConn:   wsConn,
		Session:  session,
		Uid:      uid,
		GroupNm:  "all",
		Pant:     date,
	}
	Client.Lock.Lock()
	Client.ClientList[uid] = bindInfo
	Client.Lock.Unlock()
	return
}

//.解除绑定
func (wsEvents Events) UnBind(uid string) (err error) {
	if uid == "" {
		return errors.New("uid nil")
	}
	Client.Lock.Lock()
	defer Client.Lock.Unlock()
	if _, ok := Client.ClientList[uid]; !ok {
		return
	}
	clientList := Client.ClientList
	clientList["uid"].WsConn.Close("unBind")
	delete(Client.ClientList, uid)
	return
}

//.加入通信组
func (wsEvents Events) JoinGroup(uid string, groupNm string) (err error) {
	var date int64 = time.Now().Unix()
	Client.Lock.Lock()
	defer Client.Lock.Unlock()
	if _, ok := Client.ClientList[uid]; ok {
		Client.ClientList[uid].GroupNm = groupNm
		Client.ClientList[uid].Pant = date
	} else {
		return errors.New("no uid")
	}
	return
}

//.离开通信组
func (wsEvents Events) LeaveGroup(uid string, groupNm string) (err error) {
	var date int64 = time.Now().Unix()
	Client.Lock.Lock()
	defer Client.Lock.Unlock()
	if _, ok := Client.ClientList[uid]; ok {
		Client.ClientList[uid].GroupNm = "all"
		Client.ClientList[uid].Pant = date
	} else {
		return errors.New("no uid")
	}
	return
}

//.推送消息给uid
func (wsEvents Events) SendToUid(uid string, body wspl.WsResponse) (err error) {
	if uid == "" {
		return errors.New("uid nil")
	}
	data,_ := json.Marshal(body)
	buf := string(data)
	clientInfo, err := wsEvents.GetClientInfoByUid(uid)
	if err != nil {
		return
	}
	select {
	case clientInfo.WsConn.WriteChan <- buf:
	case <-time.After(2 * time.Second):
		//.2秒无法写进数据,直接中断
		return errors.New("write timeout 5s")
	}
	return
}

//.推送给多个uid
func (wsEvents Events) SendToUidlist(uidList []string, body wspl.WsResponse) (err error) {
	if len(uidList) <= 0 {
		return errors.New("uid list len < 0")
	}
	clientList, err := wsEvents.GetClientInfoByUids(uidList)
	if err != nil {
		return
	}
	if len(clientList) <= 0 {
		return errors.New("client list len < 0")
	}
	data ,_ := json.Marshal(body)
	buf := string(data)
	for _, v := range clientList {
		//.防止线程阻塞
		select {
		case v.WsConn.WriteChan <- buf:
		case <-time.After(2 * time.Second):
			//.2秒无法写进数据,直接中断
			continue
		}
	}
	return
}

//.推送消息给socket
func (wsEvents Events) SendToSocket(ws *websocket.Conn, msg string) (err error) {
	if ws == nil || msg == "" {
		return errors.New("wspl&msg nil")
	}
	err = ws.WriteMessage(1, []byte(msg))
	if err != nil {
		return
	}
	return
}

//.给所有在线socket推送消息
func (wsEvents Events) SendToAll(body wspl.WsResponse) {
	data,_ := json.Marshal(body)
	buf := string(data)
	allClient := wsEvents.GetAllClient()
	for _, v := range allClient {
		v.WsConn.WriteChan <- buf
	}
}

//.群组推送消息
func (wsEvents Events) SendToGroup(groupNm string, msg interface{}) (err error) {
	jsonB, err := json.Marshal(msg)
	if err != nil {
		return
	}
	jsonStr := string(jsonB)
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	for _, v := range Client.ClientList {
		if v.GroupNm == groupNm {
			wsEvents.SendToSocket(v.WsConn.Conn, jsonStr)
		}
	}
	return
}

//.是否在线
func (wsEvents Events) IsOnline(uid string) (res string, err error) {
	var buf map[string]string = make(map[string]string)
	buf["online"] = "0"
	isOnline, err := wsEvents.IsOnlineByUid(uid)
	if err != nil {
		return
	}
	if isOnline == true {
		buf["online"] = "1"
	}
	clientInfo, _ := wsEvents.GetClientInfoByUid(uid)
	buf["clientId"] = clientInfo.ClientId
	buf["session"] = clientInfo.Session
	jsonB,_ := json.Marshal(buf)
	res = string(jsonB)
	return
}

//.检测uid是否在线
func (wsEvents Events) IsOnlineByUid(uid string) (online bool, err error) {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	if _, ok := Client.ClientList[uid]; !ok {
		return false, nil
	}
	return true, nil
}

//.客户端直连监测通信连接是否正常
func (wsEvents Events) GatewayCheck(wsConn *WsConn) {
	res := map[string]interface{}{
		"type": "gate",
		"msg":  "online",
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	}
	bodyB, _ := json.Marshal(res)
	//.防止线程阻塞
	select {
	case wsConn.WriteChan <- string(bodyB):
	case <-time.After(2 * time.Second):
		//.2秒无法写进数据,直接中断
		return
	}
}

//.获取群组内在线uid数量
func (wsEvents Events) GetUidCountByGroup(groupNm string) int {
	var count int = 0
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	for _, v := range Client.ClientList {
		if v.GroupNm == groupNm {
			count++
		}
	}
	return count
}

//.获取所有在线数量
func (wsEvents Events) GetAllUidCount() int {
	Client.Lock.RLock()
	defer Client.Lock.RUnlock()
	return len(Client.ClientList)
}
