package voltanet

import (
	"strconv"
	"voltanet/provider"
	"voltanet/wspl"
)

type Context struct {
	Bbo       *wspl.Bbo
	WsConn 	  *WsConn
	Code      int
	k         string
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
	c.Key(c.k)
	c.Back()
}


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
	c.SetResponse(map[string]string{"k":k}).WriteExtend(c.Bbo)
}

func (c *Context) SetResponse(data interface{}) (res *wspl.Response){
	return &wspl.Response{data}
}
