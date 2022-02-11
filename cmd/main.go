package main

import "github.com/voltaspace/volta-net"

func main()  {
	net := voltanet.NewServer("0.0.0.0:9500","json")
	net.Router("/ping", fuck)
	net.Router("/pingasdpo1o2i3u1oi2u3",fuck )
	net.Run()
}

func fuck(ctx *voltanet.Context) {
	ctx.Back()
}