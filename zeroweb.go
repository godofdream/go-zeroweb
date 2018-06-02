package zeroweb

import (
	"runtime"
	"time"

	"github.com/godofdream/fasthttp"
	"github.com/godofdream/fasthttprouter"
	"github.com/godofdream/jet"
	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Zeroweb struct {
	Config   *viper.Viper
	DB       *pgx.ConnPool
	dbConfig *pgx.ConnPoolConfig
	Log      *zerolog.Logger
	Router   *fasthttprouter.Router
	Server   *fasthttp.Server
	View     *jet.Set
}

// New takes a viper config and returns a half initialized Zeroweb.
// After adding some Routes start with it with Serve()
func New(config *viper.Viper) *Zeroweb {
	// defaults
	config.SetDefault("log.log_level", zerolog.InfoLevel.String())
	config.SetDefault("db.maxconnections", runtime.NumCPU()*4)
	config.SetDefault("http.addr", "localhost:80")
	config.SetDefault("http.MaxPendingConnections", runtime.NumCPU()*2000)
	config.SetDefault("http.read_timeout", time.Minute*1)
	config.SetDefault("http.write_timeout", time.Minute*1)
	config.SetDefault("http.max_connections_per_ip", 20)
	config.SetDefault("http.max_connections_per_connection", 20)
	config.SetDefault("http.max_keepalive_duration", time.Minute*5)
	config.SetDefault("http.max_request_body_size", 4*1024*1024)
	config.SetDefault("http.reduce_memory_consumption", false)
	config.SetDefault("http.max_concurrent_connections", 1024*1024)

	result := &Zeroweb{
		Config:   config,
		DB:       nil,
		dbConfig: nil,
		Log:      &log.Logger,
		Router:   fasthttprouter.New(),
		Server:   nil,
	}
	result.reloadLogger()
	result.reloadDB()
	return result
}
