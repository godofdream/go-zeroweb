package zeroweb

import (
	"runtime"
	"time"

	"github.com/godofdream/fasthttp"
	"github.com/godofdream/fasthttprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes httpserver, dbconnection, etc.
func (a *Zeroweb) Init() {

	// defaults
	a.Config.SetDefault("log.log_level", zerolog.InfoLevel.String())
	a.Config.SetDefault("db.maxconnections", runtime.NumCPU()*4)
	a.Config.SetDefault("http.addr", "localhost:80")
	a.Config.SetDefault("http.MaxPendingConnections", runtime.NumCPU()*2000)
	a.Config.SetDefault("http.read_timeout", time.Minute*1)
	a.Config.SetDefault("http.write_timeout", time.Minute*1)
	a.Config.SetDefault("http.max_connections_per_ip", 20)
	a.Config.SetDefault("http.max_connections_per_connection", 20)
	a.Config.SetDefault("http.max_keepalive_duration", time.Minute*5)
	a.Config.SetDefault("http.max_request_body_size", 500)
	a.Config.SetDefault("http.reduce_memory_consumption", false)
	a.Config.SetDefault("http.max_concurrent_connections", 1024*1024)

	a.Log = &log.Logger
	a.Router = fasthttprouter.New()
	a.Router.ServeFiles("/s/*filepath", "./static/")
	a.Reload()

	a.Server = &fasthttp.Server{
		Handler:              a.Router.Handler,
		Concurrency:          a.Config.GetInt("http.max_concurrent_connections"),
		Name:                 "",
		ReadTimeout:          a.Config.GetDuration("http.read_timeout"),
		WriteTimeout:         a.Config.GetDuration("http.write_timeout"),
		MaxConnsPerIP:        a.Config.GetInt("http.max_connections_per_ip"),
		MaxRequestsPerConn:   a.Config.GetInt("http.max_connections_per_connection"),
		MaxKeepaliveDuration: a.Config.GetDuration("http.max_keepalive_duration"),
		MaxRequestBodySize:   a.Config.GetInt("http.max_request_body_size"),
		ReduceMemoryUsage:    a.Config.GetBool("http.reduce_memory_consumption"),
		Logger:               a.Log,
	}

}
