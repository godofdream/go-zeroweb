package zeroweb

import (
	"runtime"
	"time"

	"github.com/go-pg/pg"
	"github.com/godofdream/fasthttp"
	"github.com/godofdream/fasthttprouter"
	"github.com/godofdream/jet"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Zeroweb struct {
	Config *viper.Viper
	DB     *pg.DB
	Log    *zerolog.Logger
	Router *fasthttprouter.Router
	Server *fasthttp.Server
	View   *jet.Set
}

// New takes a viper config and returns a half initialized Zeroweb.
// After adding some Routes start with it with Serve()
func New(config *viper.Viper) *Zeroweb {
	// defaults
	config.SetDefault("log.LogLevel", zerolog.InfoLevel.String())
	config.SetDefault("log.LevelFieldName", "level")
	config.SetDefault("log.MessageFieldName", "message")
	config.SetDefault("log.TimestampFieldName", "time")

	config.SetDefault("db.Enabled", true)

	config.SetDefault("db.Network", "tcp")
	config.SetDefault("db.Addr", "127.0.0.1:5432")
	config.SetDefault("db.MaxRetries", 0)
	config.SetDefault("db.RetryStatementTimeout", false)
	config.SetDefault("db.MinRetryBackoff", 250*time.Millisecond)
	config.SetDefault("db.MaxRetryBackoff", 4*time.Second)
	config.SetDefault("db.DialTimeout", 5*time.Second)
	config.SetDefault("db.ReadTimeout", 0)
	config.SetDefault("db.WriteTimeout", 0)
	config.SetDefault("db.PoolSize", 10*runtime.NumCPU())
	config.SetDefault("IdleCheckFrequency", 1*time.Minute)

	config.SetDefault("http.Name", "") // dropping servername in http as it uses unnecessary bytes and tells attackers about the system
	config.SetDefault("http.Concurrency", 1048576)
	config.SetDefault("http.DisableKeepalive", true)
	config.SetDefault("http.ReadBufferSize", 8192)
	config.SetDefault("http.WriteBufferSize", 8192)
	config.SetDefault("http.ReadTimeout", 0)
	config.SetDefault("http.WriteTimeout", 0)
	config.SetDefault("http.MaxConnsPerIP", 0)
	config.SetDefault("http.MaxRequestsPerConn", 0)
	config.SetDefault("http.MaxKeepaliveDuration", 0)
	config.SetDefault("http.MaxRequestBodySize", 4194304)
	config.SetDefault("http.ReduceMemoryUsage", false)
	config.SetDefault("http.GetOnly", false)
	config.SetDefault("http.LogAllErrors", false)
	config.SetDefault("http.DisableHeaderNamesNormalizing", false)
	config.SetDefault("http.NoDefaultServerHeader", true)

	config.SetDefault("http.Addr", "localhost:8080")

	config.SetDefault("templates.Folder", "./templates")

	// config.SetDefault("templates.folder", "./templates")
	config.SetDefault("static.StaticFolder", "./static")
	config.SetDefault("static.CSSFolder", "./static/css")
	config.SetDefault("static.JSFolder", "./static/js")
	config.SetDefault("static.FontsFolder", "./static/fonts")
	// config.SetDefault("cache.assetcache_folder", "./cache/assets")
	// config.SetDefault("cache.optimize_assets", false)
	result := &Zeroweb{
		Config: config,
		DB:     nil,
		Log:    &log.Logger,
		Router: fasthttprouter.New(),
		Server: nil,
	}
	result.reloadLogger()
	result.reloadDB()
	result.reloadHTTP()
	return result
}
