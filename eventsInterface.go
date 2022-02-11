package voltanet

import (
	"github.com/gorilla/websocket"
	"github.com/voltaspace/volta-net/wspl"
)

type EventsInterface interface {
	//.查询客户端信息
	GetClientInfo(clientId string) (res ClientInfo, err error)
	//.通过uid获取绑定信息
	GetClientInfoByUid(uid string) (res ClientInfo, err error)
	//.获取所有客户端
	GetAllClient() (clients []ClientInfo)
	//.同步uidlist获取对应通道
	GetClientInfoByUids(uidlist []string) (clientList []ClientInfo, err error)
	//.通过ws指针查找client信息
	GetClientInfoBySocket(ws *websocket.Conn) (clinet ClientInfo,err error)
	//.强制断开uid
	CloseSocketByUid(uid string) (err error)
	//.client绑定uid
	Bind(session string, uid string, wsConn *WsConn) (clientId string, err error)
	//.解除绑定
	UnBind(uid string) (err error)
	//.加入通信组
	JoinGroup(uid string, groupNm string) (err error)
	//.离开通信组
	LeaveGroup(uid string, groupNm string) (err error)
	//.推送消息给uid
	SendToUid(uid string, body wspl.WsResponse) (err error)
	//.推送给多个uid
	SendToUidlist(uidList []string, body wspl.WsResponse) (err error)
	//.推送消息给socket
	SendToSocket(ws *websocket.Conn, msg string) (err error)
	//.给所有在线socket推送消息
	SendToAll(body wspl.WsResponse)
	//.群组推送消息
	SendToGroup(groupNm string, msg interface{}) (err error)
	//.是否在线
	IsOnline(uid string) (res string, err error)
	//.检测uid是否在线
	IsOnlineByUid(uid string) (online bool, err error)
	//.客户端直连监测通信连接是否正常
	GatewayCheck(wsConn *WsConn)
	//.获取群组内在线uid数量
	GetUidCountByGroup(groupNm string) int
	//.获取所有在线数量
	GetAllUidCount() int
}
