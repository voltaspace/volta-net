package wspl

import (
	"encoding/json"
	"errors"
	"github.com/voltaspace/volta-net/constant/cmd"
	"github.com/voltaspace/volta-net/show"
	"strings"
)

func (wsRequest *WsRequest) Render(wspl Wspl) (err error) {
	var bbo Bbo
	bbo.Request = wsRequest
	_, err = wspl.Analysis(&bbo)
	if err != nil {
		return
	}
	err = wspl.Write(&bbo)
	if err != nil {
		return
	}
	err = wspl.WriteHeader(&bbo)
	if err != nil {
		return
	}
	err = wspl.WriteEnd(&bbo)
	if err != nil {
		return
	}
	err = wspl.WriteExtend(&bbo)
	if err != nil {
		return
	}
	return
}

type Request struct {
	Data interface{}
}

func (write *Request) Analysis(bbo *Bbo) (buf string, err error) {
	body := write.Data.(string)
	if len(body) <= 15 || body[0:15] != cmd.STRIKE {
		err = errors.New("[volta-gateway]:is not gateway protocol")
		return
	}
	deviationStr := body[cmd.DEVIATION_LEN:]
	if deviationStr == "" {
		err = errors.New("[volta-gateway]:deviationLen nil")
		return
	}
	//.[0]header [1]body [2]end [3]extend
	bufs := strings.Split(deviationStr, cmd.GATEWAY_PADDING)
	newBufs := show.RemoveEmptySlice(bufs)
	write.Data = newBufs
	return body, nil
}

func (req *Request) WriteHeader(bbo *Bbo) (err error) {
	bufs := req.Data.([]string)
	err = json.Unmarshal([]byte(bufs[0]), &bbo.Request.Header)
	if err != nil {
		return
	}
	return
}

func (req *Request) Write(bbo *Bbo) (err error) {
	bufs := req.Data.([]string)
	bbo.Request.Data = bufs[1]
	if bbo.Request.Data == ""{
		bbo.Request.Param = map[string]interface{}{}
		return
	}
	err = json.Unmarshal([]byte(bufs[1]), &bbo.Request.Param)
	if err != nil {
		return
	}
	return
}

func (req *Request) WriteEnd(bbo *Bbo) (err error) {
	bufs := req.Data.([]string)
	err = json.Unmarshal([]byte(bufs[2]), &bbo.Request.End)
	if err != nil {
		return
	}
	return
}

func (req *Request) WriteExtend(bbo *Bbo) (err error) {
	bufs := req.Data.([]string)
	err = json.Unmarshal([]byte(bufs[3]), &bbo.Request.Extend)
	if err != nil {
		return
	}
	return
}

func (wsReq *WsRequest) GetBody() string {
	return wsReq.Data
}

func (wsReq *WsRequest) GetParam(key string) interface{}{
	return wsReq.Param[key]
}

func (WsReq *WsRequest) GetParams() map[string]interface{}{
	return WsReq.Param
}

func (wsReq *WsRequest) GetHeader() Header {
	return wsReq.GetHeader()
}
