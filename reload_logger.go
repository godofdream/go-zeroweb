package zeroweb

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *Zeroweb) ReloadLogger() error {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = ""
	loglevelstr := a.Config.GetString("log.log_level")
	level, err := zerolog.ParseLevel(loglevelstr)
	if err != nil {
		log.Fatal().Err(err).Str("loglevel", loglevelstr).Msg("Bad Loglevel")
		return err
	} else {
		zerolog.SetGlobalLevel(level)
		return nil
	}
}
