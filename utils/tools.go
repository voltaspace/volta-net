package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//.转换为json格式
func JsonEncode(data interface{})(string){
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}

func JsonDecode(t interface{},data string) interface{}{
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return t
	}
	return t
}

func GetDate() int64{
	return time.Now().Unix()
}

func GetTime() string{
	var date string = time.Now().Format("2006-01-02 15:04:05")
	return date
}

func GetTimeStamp() string{
	var date string = time.Now().Format("20060102150405")
	return date
}

func PrintStruct(info interface{}){
	fmt.Printf("%+v\n",info)
}

func Echo(data interface{})  {
	fmt.Println(fmt.Sprintf("%+v",data))
}
//.获取随机字符串
func  GetRandomString(l int) string {
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


func Rank(str string,cnt int) string{
	var k1 string
	for i:=1;i<cnt-len(str);i++{
		k1 += " "
	}
	return k1
}

func VoltaPrint(strs ...string){
	var printStr string = "[volta] CMD:"
	for k,v := range strs{
		if k == 1 {
			printStr += v + "" + Rank(v,40)
			continue
		}
		//if k == 2 {
		//	printStr += "---> "
		//}
		//printStr += v + "" + Rank(v,25)
	}
	fmt.Println(printStr)

}
