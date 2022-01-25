package cmd

import "voltanet"

func main()  {
	net := voltanet.NewVoltaNet("0.0.0.0",19999,"json")
	net.Router("/ping", func(ctx *voltanet.Context) {
		ctx.Back()
	})
	net.Run()
}