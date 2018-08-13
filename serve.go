package zeroweb

import (
	"github.com/go-siris/tcplisten"
	"github.com/rs/zerolog/log"
)

func (zeroweb *Zeroweb) Serve() error {
	zeroweb.Reload()
	tcpconf := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
		Backlog:     zeroweb.Config.GetInt("http.max_pending_connections"),
	}
	network := zeroweb.Config.GetString("http.network")
	addr := zeroweb.Config.GetString("http.addr")
	ln, err := tcpconf.NewListener(network, addr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed starting Listener")
		return err
	}

	log.Info().Str("network", network).Str("addr", addr).Msg("Starting Webserver")
	if err = zeroweb.Server.Serve(ln); err != nil {
		log.Fatal().Err(err).Msg("Failed serving incoming connections")
		return err
	}

	return nil
}
