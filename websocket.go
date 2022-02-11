package voltanet

import (
	"fmt"
	"github.com/google/wire"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Options struct {
	Addr string
	CtxType string
}

var OptionsProvider = wire.NewSet(NewOptions)

func NewOptions() *Options{
	return &Options{}
}


type Websocket struct {
	Signal *Signal
	Events *Events
	WsHeartBeat *WsHeartBeat
	Options *Options
}

var WebSocketProvider = wire.NewSet(NewWebsocket)

func NewWebsocket(signal *Signal,events *Events,wsHeartBeat *WsHeartBeat, options *Options) *Websocket{
	return &Websocket{
		signal,
		events,
		wsHeartBeat,
		options,
	}
}

//.开启socket服务
func (ws *Websocket) Run() {
	http.HandleFunc("/", ws.GatewayManager)
	fmt.Printf("[volta-net] websocket service start (addr %s)",ws.Options.Addr)

	err := http.ListenAndServe(ws.Options.Addr, nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %s", err.Error())
	}
}

func (ws *Websocket) SetOptions(options *Options){
	ws.Options = options
}

func (ws *Websocket) GatewayManager(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET wspl to a websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		go ws.Alloter(conn)
	}
}

func (ws *Websocket) Alloter(conn *websocket.Conn) {
	defer EndStack("[volta-net]:socket master process closed")
	wsConn := &WsConn{
		Conn:    conn,
		WriteChan: make(chan string, 1000),
		Signal : ws.Signal,
		WsHeartBeat: ws.WsHeartBeat,
		CtxType: ws.Options.CtxType,
		Closed:    false,
	}
	ws.Signal.OnConnect(conn, "guest")
	go wsConn.readProcess()
	go wsConn.writeProcess()
	for {
		//.主线程阻塞 5秒检测一次是否断连
		if wsConn.Closed == true {
			break
		}
		time.Sleep(5 * time.Second)
	}
	//.主线程退出前关闭socket链接,这里会出现空指针错误，defer已经忽略错误不会导致crash
	conn.Close()
}

