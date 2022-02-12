package voltanet

type TransServer interface {
	Run()
	SetOptions(options *Options)
}

type App struct {
	Server []TransServer
}

func NewApp(transServer ...TransServer) *App{
	return &App{transServer}
}

func (app *App) SetOptions(options *Options){
	for k,_ := range app.Server {
		app.Server[k].SetOptions(options)
	}
}

func (app *App) Run(){
	for _,s := range app.Server {
		s.Run()
	}
	select {

	}
}
