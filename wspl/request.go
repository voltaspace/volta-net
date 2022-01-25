package wspl

import (
	"encoding/json"
	"errors"
	"strings"
)

var strike = "volta-http/wspl"
var deviationLen = len(strike) + 1

const (
	GATEWAY_PADDING string = "<!#Volta>" //.ws通信数据包填充字符
)

func (wsRequest *WsRequest) Render(wspl Wspl) (err error) {
	var bbo Bbo
	bbo.Request = wsRequest
	_ ,err = wspl.Analysis(&bbo)
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

func (write *Request) Analysis(bbo *Bbo) (buf string,err error) {
	body := write.Data.(string)
	if len(body) <= 15 || body[0:15] != strike {
		err = errors.New("[volta-gateway]:is not gateway protocol")
		return
	}
	deviationStr := body[deviationLen:]
	if deviationStr == "" {
		err = errors.New("[volta-gateway]:deviationLen nil")
		return
	}
	//.[0]header [1]body [2]end [3]extend
	bufs := strings.Split(deviationStr, GATEWAY_PADDING)
	write.Data = bufs
	return body,nil
}

func (write *Request) WriteHeader(bbo *Bbo) (err error) {
	bufs := write.Data.([]string)
	err = json.Unmarshal([]byte(bufs[0]), &bbo.Request.Header)
	if err != nil {
		return
	}
	return
}

func (write *Request) Write(bbo *Bbo) (err error) {
	bufs := write.Data.([]string)
	err = json.Unmarshal([]byte(bufs[1]), &bbo.Request.Data)
	if err != nil {
		return
	}
	return
}

func (write *Request) WriteEnd(bbo *Bbo) (err error) {
	bufs := write.Data.([]string)
	err = json.Unmarshal([]byte(bufs[2]),&bbo.Request.End)
	if err != nil {
		return
	}
	return
}

func (write *Request) WriteExtend(bbo *Bbo) (err error) {
	bufs := write.Data.([]string)
	err = json.Unmarshal([]byte(bufs[3]), &bbo.Request.Extend)
	if err != nil {
		return
	}
	return
}

