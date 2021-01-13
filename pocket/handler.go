package pocket

import "github.com/valyala/fasthttp"

type FastHttpHandler func(ctx *fasthttp.RequestCtx) interface{}

func (h FastHttpHandler) Handler(ctx *fasthttp.RequestCtx) {
	value := h(ctx)
	ctx.SetContentTypeBytes([]byte("application/json"))
	var resp *Response

	switch v := value.(type) {
	case error:
		resp = NewResponseFromError(v)
	case *Response:
		resp = v
	case nil:
		resp = NewResponse(204, nil)
	default:
		resp = NewResponse(200, v)
	}
	resp.ResponseWithFastHttpCtx(ctx)
}
