package voltanet

import (
	"github.com/voltaspace/volta-net/provider"
	"github.com/voltaspace/volta-net/wspl"
	"strconv"
)

type Context struct {
	Bbo       *wspl.Bbo
	WsConn 	  *WsConn
	Code      int
	K         string
}

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func (c *Context) JSON(obj interface{}) {
	c.Render(wspl.Success, provider.JSON{obj})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (c *Context) ProtoBuf(obj interface{}) {
	c.Render(wspl.Success, provider.ProtoBuf{obj})
}

// Render writes the response headers and calls render.Render to render data.
func (c *Context) Render(code int, p provider.Render) {
	c.Status(code)
	err := p.Render(c.Bbo)
	if err != nil {
		panic(err)
	}
	p.WriteContentType(c.Bbo)
	c.Key(c.K)
	c.Back()
}


func (c *Context) Back() {
	res := wspl.Response{}
	buf,err := res.Analysis(c.Bbo)
	if err != nil {
		return
	}
	c.WsConn.SendSelf(buf,true)
}

func (c *Context) Status(code int) {
	c.SetResponse(map[string]string{"status":strconv.FormatInt(int64(code),10)}).WriteExtend(c.Bbo)
}

func (c *Context) Key(k string) {
	c.SetResponse(map[string]string{"seq":k}).WriteHeader(c.Bbo)
}

func (c *Context) Request() *wspl.WsRequest{
	return c.Bbo.Request
}

func (c *Context) Response() *wspl.WsResponse {
	return c.Bbo.Response
}

func (c *Context) SetResponse(data interface{}) (res *wspl.Response){
	return &wspl.Response{data}
}
