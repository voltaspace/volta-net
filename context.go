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

// NewContext Create a new context to push messages only.
// This context has only response information and no request information.
func NewContext(ctxType string) *Context{
	var context Context
	context.Bbo = &wspl.Bbo{
		Request: nil,
		Response: &wspl.WsResponse{
			Ctx:    ctxType,
			Header: make(map[string]string),
			Data:   make([]byte, 0),
			Extend: make(map[string]string),
		},
	}
	context.K = ""
	context.WsConn = nil
	return &context
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

// GetDefault
func (c *Context) GetDefault() (buf string){
	res := wspl.Response{}
	buf,_ = res.Analysis(c.Bbo)
	return
}

// BuildJSON
func (c *Context) BuildJSON(obj interface{}) (buf string){
	return c.RenderStr(wspl.Success, provider.JSON{obj})
}

// BuildProtoBuf
func (c *Context) BuildProtoBuf(obj interface{}) (buf string){
	return c.RenderStr(wspl.Success, provider.ProtoBuf{obj})
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

// Render writes the response headers and calls render.Render to render data & return to string.
func (c *Context) RenderStr(code int, p provider.Render) (buf string){
	c.Status(code)
	err := p.Render(c.Bbo)
	if err != nil {
		panic(err)
	}
	p.WriteContentType(c.Bbo)
	c.Key(c.K)
	res := wspl.Response{}
	buf,_ = res.Analysis(c.Bbo)
	return
}

// Back
func (c *Context) Back() {
	res := wspl.Response{}
	buf,err := res.Analysis(c.Bbo)
	if err != nil {
		return
	}
	c.WsConn.SendSelf(buf)
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
