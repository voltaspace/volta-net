package voltanet

import (
	"github.com/voltaspace/volta-net/container"
)

type HandlerFunc func(ctx *Context)

func newApp(wsHeartBeat *WsHeartBeat,ws *Websocket) *App {
	return NewApp(
		ws,
		wsHeartBeat,
	)
}

type VoltaNet struct {
	Signal  Signal
	CtxType string
	Events  *Events
	Options *Options
}

func NewVoltaNet(host string, port int, ctxType string) *VoltaNet {
	Register()
	return &VoltaNet{
		Signal:  Signal{},
		CtxType: ctxType,
		Events:  NewEvents(),
		Options: &Options{host, port, ctxType},
	}
}

func (g *VoltaNet) Run() {
	app,_,err := initApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}

func (g *VoltaNet) Router(relativePath string, handlers HandlerFunc) {
	container.App.Bind(relativePath, handlers)
}

func (g *VoltaNet) Use() {

}

