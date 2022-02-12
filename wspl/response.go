package wspl

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/voltaspace/volta-net/protoc"
	"time"
)

type Response struct {
	Data interface{}
}

const (
	Success = 200
	Error   = -1
)

func (write *Response) Analysis(bbo *Bbo) (buff string, err error) {
	bodyMainStr := string(bbo.Response.Data)
	header := bbo.Response.Header
	extend := bbo.Response.Extend
	var resB []byte
	if bbo.Response.Ctx == "json" {
		resB,err = JSON(header,bodyMainStr,extend)
	}else if bbo.Response.Ctx == "protobuf"{
		resB,err = PROTOBUF(header,bodyMainStr,extend)
	}
	if err != nil {
		return
	}
	return string(resB), nil
}

func JSON(header map[string]string, body string, extend map[string]string) (resB []byte,err error){
	//.使用Json作为基本传输协议
	communication := map[string]interface{}{
		"header": header,
		"apply":  body,
		"extend": extend,
		"ghoul":  int64(len(header)),
		"bee":    int64(len(body)),
		"elen":   int64(len(extend)),
		"t":      time.Now().Unix(),
	}
	resB, err = json.Marshal(communication)
	if err != nil {
		return
	}
	return
}
func PROTOBUF(header map[string]string, body string, extend map[string]string)  (resB []byte,err error){
	//.使用protobuf作为基本传输协议
	communication := &protoc.CommunicationMain{
		Header: header,
		Apply:  body,
		Extend: extend,
		Ghoul:  int64(len(header)),
		Bee:    int64(len(body)),
		Elen:   int64(len(extend)),
		T:      time.Now().Unix(),
	}
	resB, err = proto.Marshal(communication)
	if err != nil {
		return
	}
	return
}

func (write *Response) Write(bbo *Bbo) (err error) {
	bbo.Response.Data = write.Data.([]byte)
	return
}

func (write *Response) WriteHeader(bbo *Bbo) (err error) {
	header := write.Data.(map[string]string)
	for k, v := range header {
		bbo.Response.Header[k] = v
	}
	return
}

func (write *Response) WriteEnd(bbo *Bbo) (err error) {
	return
}

func (write *Response) WriteExtend(bbo *Bbo) (err error) {
	extend := write.Data.(map[string]string)
	for k, v := range extend {
		bbo.Response.Extend[k] = v
	}
	return
}
