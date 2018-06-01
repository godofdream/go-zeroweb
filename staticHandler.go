package zeroweb

import "github.com/godofdream/fasthttp"

func (a *Zeroweb) staticHandler(ctx *fasthttp.RequestCtx) {
	ctx.SendFileBytes(ctx.Path())
}
