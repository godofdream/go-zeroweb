package pprofhandler

import (
	"net/http/pprof"
	"strings"

	"github.com/godofdream/fasthttp"
	"github.com/godofdream/fasthttp/fasthttpadaptor"
)

// PprofHandler serves server runtime profiling data in the format expected by the pprof visualization tool.
//
// See https://golang.org/pkg/net/http/pprof/ for details.
func PprofHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "text/html")
	if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/cmdline") {
		fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/profile") {
		fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/symbol") {
		fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/trace") {
		fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)(ctx)
	} else {
		fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)(ctx)
	}
}
