package pocket

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type Response struct {
	Code int
	Data interface{}
}

func NewResponse(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}

func (r *Response) ResponseWithFastHttpCtx(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(r.Code)
	if r.Data == nil {
		return
	}
	err := json.NewEncoder(ctx).Encode(r.Data)
	if err == nil {
		return
	}
	ctx.Error(err.Error(), r.Code)
}
