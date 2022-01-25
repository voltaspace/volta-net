package voltanet

type TransServer interface {
	Run()
}

type App struct {
	Server []TransServer
}

func NewApp(transServer ...TransServer) *App{
	return &App{transServer}
}

func (app *App) Run(){
	for _,s := range app.Server {
		s.Run()
	}
}
