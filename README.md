# volta-net
volta-net
单元启动 cmd/main.go
# 使用教程
net := voltanet.NewVoltaNet("0.0.0.0",19999,"json")
net.Router("/ping", func(ctx *voltanet.Context) {
		ctx.Back()
})
net.Run()
