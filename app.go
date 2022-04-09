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

// Run Start all services. When suspend is equal to 0, it will be executed asynchronously.
// When suspend is equal to 1, it will be suspended in the background
func (app *app) Run(suspend int){
	App = app
	for _,s := range app.Server {
		s.Run()
	}
	if suspend == 1 {
		select {}
	}
}
