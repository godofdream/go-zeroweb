package zeroweb

import "github.com/valyala/fasthttp"

func (a *Zeroweb) staticHandler(ctx *fasthttp.RequestCtx) {
	ctx.SendFileBytes(ctx.Path())
}
