package voltanet

import (
	"github.com/voltaspace/volta-net/container"
)

type HandlerFunc func(ctx *Context)

func newApp(wsHeartBeat *WsHeartBeat,ws *Websocket) *app {
	return NewApp(
		ws,
		wsHeartBeat,
	)
}

type Server struct {
	Signal  Signal
	CtxType string
	Events  *Events
	Options *Options
}

func NewServer(addr string, ctxType string) *Server {
	Register()
	return &Server{
		Signal:  Signal{},
		CtxType: ctxType,
		Events:  NewEvents(),
		Options: &Options{addr, ctxType},
	}
}

func (g *Server) Run() {
	app,_,err := initApp()
	if err != nil {
		panic(err)
	}
	app.SetOptions(g.Options)
	app.SetEvents(g.Events)
	app.Run()
}

func (g *Server) Router(relativePath string, handlers HandlerFunc) {
	container.App.Bind(relativePath, handlers)
}

func (g *Server) Use() {

}

