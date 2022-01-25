package utils

import (
	"fmt"
	"runtime/debug"
)

func EndStack(appName string) {
	if err := recover(); err != nil {
		//将客户端的这次请求头、主体等信息+程序的堆栈信息
		msg := map[string]interface{}{
			"error": err,                   //真正的错误信息
			"wspl":  appName,               //连接句柄信息
			"stack": string(debug.Stack()), //此刻程序的堆栈信息
		}
		fmt.Println(msg)
	}
}