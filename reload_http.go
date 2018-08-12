package zeroweb

import (
	"github.com/godofdream/fasthttp"
	"github.com/godofdream/jet"
)

func (a *Zeroweb) reloadHTTP() error {
	a.View = jet.NewHTMLSet("./templates")

	a.Server = &fasthttp.Server{
		Handler:              a.Router.Handler,
		Concurrency:          a.Config.GetInt("http.max_concurrent_connections"),
		Name:                 "", // dropping servername in http as it uses unnecessary bytes and tells attackers about the system
		ReadTimeout:          a.Config.GetDuration("http.read_timeout"),
		WriteTimeout:         a.Config.GetDuration("http.write_timeout"),
		MaxConnsPerIP:        a.Config.GetInt("http.max_connections_per_ip"),
		MaxRequestsPerConn:   a.Config.GetInt("http.max_connections_per_connection"),
		MaxKeepaliveDuration: a.Config.GetDuration("http.max_keepalive_duration"),
		MaxRequestBodySize:   a.Config.GetInt("http.max_request_body_size"),
		ReduceMemoryUsage:    a.Config.GetBool("http.reduce_memory_consumption"),
		Logger:               a.Log,
	}

	return nil
}
