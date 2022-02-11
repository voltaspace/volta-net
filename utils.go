package voltanet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"runtime/debug"
	"strconv"
	"time"
)

type H map[string]interface{}

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

func GetTime() string{
	var date string = time.Now().Format("2006-01-02 15:04:05")
	return date
}

func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Log(info string){
	fmt.Println(GetTime() + " info:"+ info)
}

func RandInt(min int, max int) int {
	if min >= max || min==0 || max==0{
		return max
	}
	return rand.Intn(max-min)+min
}

func ParseFloat32(value string) float32{
	res,_ := strconv.ParseFloat(value,32)
	return float32(res)
}

func ParseInt(value string) int{
	res,_ := strconv.Atoi(value)
	return int(res)
}