# volta-net
volta-net
单元启动 cmd/main.go
# 使用教程
~~~~
net := voltanet.NewVoltaNet("0.0.0.0",19999,"json")  
net.Router("/ping", func(ctx *voltanet.Context) {  
	ctx.Back()  
})
net.Run()
~~~~
# 功能请求
~~~~
1.ping 心跳
2.gate 模拟请求
3.protocol 协议模板
~~~~
# 默认协议
# Request
~~~~
偏移码：volta-http/wspl:
填充码：#volta-net#
头部信息：{"ExtendHeader":{"auth":"test"},"seq":"WTZ3QOVQFC6XX4RVIH8GNPB6YDSLJ0JO","session":"session","method":"ping","contentType":"application/json","protocol":"volta"}
包体信息：{"body":"test"}
尾部信息：{"bodyLength":15}
扩展信息：{"extend":"extend"}
完整请求体：volta-http/wspl:#volta-net#{"ExtendHeader":{"auth":"test"},"seq":"WTZ3QOVQFC6XX4RVIH8GNPB6YDSLJ0JO","session":"session","method":"ping","contentType":"application/json","protocol":"volta"}#volta-net#{"body":"test"}#volta-net#{"bodyLength":15}#volta-net#{"extend":"extend"}#volta-net#
~~~~
# Response
~~~~
头部信息："header":{"contentType":"application/json","seq":"WTZ3QOVQFC6XX4RVIH8GNPB6YDSLJ0JO"}
包体信息："apply":"123"
{"apply":"123","bee":5,"elen":1,"extend":{"status":"200"},"ghoul":2,"header":{"contentType":"application/json","seq":"WTZ3QOVQFC6XX4RVIH8GNPB6YDSLJ0JO"},"t":1644650383}
~~~~

