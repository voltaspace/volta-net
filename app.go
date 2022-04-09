package voltanet

type transServer interface {
	Run()
	SetOptions(options *Options)
}

var App *app

type app struct {
	Server []transServer
	Events EventsInterface
}

func NewApp(transServer ...transServer) *app{
	return &app{Server:transServer}
}

func (app *app) SetOptions(options *Options){
	for k,_ := range app.Server {
		app.Server[k].SetOptions(options)
	}
}

func (app *app) SetEvents(events EventsInterface){
	app.Events = events
}

func (app *app) Run(){
	App = app
	for _,s := range app.Server {
		s.Run()
	}
	select {

	}
}
