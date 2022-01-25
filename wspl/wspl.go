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
	ExtendHeader map[string]string
	Seq         string `json:"seq"`
	Session     string `json:"session"`
	Method      string `json:"method"`
	ContentType string `json:"contentType"`
	Protocol    string `json:"protocol"`
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

