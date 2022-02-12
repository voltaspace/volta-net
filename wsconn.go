package voltanet

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/voltaspace/volta-net/container"
	"github.com/voltaspace/volta-net/wspl"
	"reflect"
	"runtime/debug"
	"time"
)

type WsConn struct {
	Conn        *websocket.Conn
	WriteChan   chan string
	Signal      *Signal
	WsHeartBeat *WsHeartBeat
	CtxType     string
	Closed      bool
}

//.只接收client发来的部分信息
func (wsConn *WsConn) readProcess() {
	defer wsConn.Close("[volta-gateway]:read disconnect")
	for {
		_, message, err := wsConn.Conn.ReadMessage()
		if err != nil {
			return
		}
		wsConn.gateWayControl(string(message))
	}
}

//.串行向当前client 推送消息
func (wsConn *WsConn) writeProcess() {
	defer wsConn.Close("[volta-gateway]:write disconnect")
	for {
		select {
		case message := <-wsConn.WriteChan:
			err := wsConn.Conn.WriteMessage(1, []byte(message))
			if err != nil {
				return
			}
		case <-time.After(15 * time.Second):
			//.15秒没有发消息请求，检测一次是否已经断连
			if wsConn.Closed == true {
				return
			}
		}
	}
}

//.向自己发消息
func (wsConn *WsConn) SendSelf(msg interface{},isJson bool) (err error) {
	var sendMsg string
	if isJson {
		jsonB, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		sendMsg = string(jsonB)
	}else{
		sendMsg = msg.(string)
	}
	select {
	case wsConn.WriteChan <- sendMsg:
	case <-time.After(2 * time.Second):
		//.2秒无法写进数据,直接中断
		return errors.New("write timeout 2s")
	}
	return
}

//.发送关闭socket链接信号
func (wsConn *WsConn) Close(appName string) {
	wsConn.Closed = true
	if err := recover(); err != nil {
		//将客户端的这次请求头、主体等信息+程序的堆栈信息
		msg := map[string]interface{}{
			"error": err,                   //真正的错误信息
			"wspl":  appName,               //连接句柄信息
			"stack": string(debug.Stack()), //此刻程序的堆栈信息
		}
		fmt.Println(msg)
	}
}

//.Gateway消息控制器
func (wsConn *WsConn) gateWayControl(message string) {
	wsConn.Signal.OnMessage(wsConn.Conn, message)
	if message == "gate" {
		wsConn.gate()
	}
	if message == "ping" {
		wsConn.WsHeartBeat.HeartBeatUpdate(wsConn.Conn)
	}
	if message == "protocol" {
		wsConn.protocol()
	}
	//.解析数据包，生成request数据
	request, err := wsConn.gatewayDecode(message)
	if err != nil {
		return
	}
	//.路由转发
	err2 := wsConn.route(request)
	if err2 != nil {
		wsConn.SendSelf(err2.Error(),false)
		return
	}
}

//.ws通信自定协议解析
func (wsConn *WsConn) gatewayDecode(message string) (req *wspl.WsRequest, err error) {
	var wsReq *wspl.WsRequest = &wspl.WsRequest{}
	err = wsReq.Render(&wspl.Request{message})
	if err != nil {
		return
	}
	req = wsReq
	return
}

//.ws任务分配路由
func (wsConn *WsConn) route(req *wspl.WsRequest) (err error) {
	relay, err := container.App.Make(req.Header.Method)
	if err != nil {
		return errors.New("header method not found")
	}
	//.创建relay反射对象 这里多例，为保证并发安全
	//refV := reflect.New(relay.ConType)
	//method := refV.MethodByName(relay.Method)
	//if !method.IsValid() {
	//	return
	//}
	var context Context
	context.Bbo = &wspl.Bbo{
		Request: req,
		Response: &wspl.WsResponse{
			Ctx:    wsConn.CtxType,
			Header: make(map[string]string),
			Data:   make([]byte, 0),
			Extend: make(map[string]string),
		},
	}
	context.K = req.Header.Seq
	context.WsConn = wsConn
	paramList := make([]reflect.Value, 1)
	paramList[0] = reflect.ValueOf(&context)
	relay.ConValue.Call(paramList)
	return
}

//. 响应客户端
func (wsConn *WsConn) gate() {
	res := map[string]interface{}{
		"type": "gate",
		"msg":  "online",
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	}
	wsConn.SendSelf(res,true)
}

//. 返回协议模板
func (wsConn *WsConn) protocol() {
	module := ProtocolModule()
	wsConn.SendSelf(module,false)
}
