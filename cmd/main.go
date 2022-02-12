package main

import "github.com/voltaspace/volta-net"

func main()  {
	net := voltanet.NewServer("0.0.0.0:9500","json")
	net.Router("ping", fuck)
	net.Run()
}

func fuck(ctx *voltanet.Context) {
	ctx.JSON(voltanet.H{
		"data":123,
	})
}