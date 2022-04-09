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

// Run all core services synchronously
func (g *Server) Run() {
	g.Init().Run(1)
}

// Run all core services asynchronously
func (g *Server) Defualt() {
	g.Init().Run(0)
}

// Init Initialization, dependency injection
func (g *Server) Init() (newApp *app) {
	newApp,_,err := initApp()
	if err != nil {
		panic(err)
	}
	newApp.SetOptions(g.Options)
	newApp.SetEvents(g.Events)
	return
}

func (g *Server) Router(relativePath string, handlers HandlerFunc) {
	container.App.Bind(relativePath, handlers)
}

func (g *Server) Use() {

}

