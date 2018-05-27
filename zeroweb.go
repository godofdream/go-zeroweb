package zeroweb

import (
	"github.com/godofdream/fasthttp"
	"github.com/godofdream/fasthttprouter"
	"github.com/godofdream/jet"
	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
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
