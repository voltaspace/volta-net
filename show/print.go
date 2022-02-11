package show

import "fmt"

func Rank(str string,cnt int) string{
	var k1 string
	for i:=1;i<cnt-len(str);i++{
		k1 += " "
	}
	return k1
}

func VoltaPrint(strs ...string){
	var printStr string = "[volta-net] hybrid:"
	for k,v := range strs{
		if k == 0 {
			printStr += v + "" + Rank(v,40)
			continue
		}
		printStr += v
	}
	fmt.Println(printStr)

}

func PrintStruct(info interface{}){
	fmt.Printf("%+v\n",info)
}

func Echo(data interface{})  {
	fmt.Println(fmt.Sprintf("%+v",data))
}
