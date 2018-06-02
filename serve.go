package zeroweb

import (
	"github.com/go-siris/tcplisten"
	"github.com/rs/zerolog/log"
)

func (a *Zeroweb) Serve() error {
	a.Reload()
	tcpconf := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
		Backlog:     a.Config.GetInt("http.max_pending_connections"),
	}
	ln, err := tcpconf.NewListener("tcp", a.Config.GetString("http.addr"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed starting Listener")
		return err
	}

	log.Info().Msg("Starting Webserver")
	if err = a.Server.Serve(ln); err != nil {
		log.Fatal().Err(err).Msg("Failed serving incoming connections")
		return err
	}

	return nil
}
