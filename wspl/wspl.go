package wspl

type Wspl interface {
	Analysis(bbo *Bbo) (buf string,err error)
	WriteHeader(bbo *Bbo) (err error)
	Write(bbo *Bbo) (err error)
	WriteEnd(bbo *Bbo) (err error)
	WriteExtend(bbo *Bbo) (err error)
}

type Bbo struct {
	Request *WsRequest
	Response *WsResponse
}


type WsRequest struct {
	Header Header
	Data   interface{}
	End    End
	Extend map[string]string
}

type Header struct {
	ContentType string `json:"contentType"`
	Method      string `json:"method"`
	Seq         string `json:"seq"`
	Session     string `json:"session"`
	Protocol    string `json:"protocol"`
	ExtendHeader map[string]string `json:"extendHeader"`
}

//.ws包尾部
type End struct {
	BodyLength int `json:"bodyLength"`
}

type WsResponse struct {
	Ctx string
	Header map[string]string
	Data []byte
	Extend map[string]string
}

