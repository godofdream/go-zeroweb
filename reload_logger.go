package zeroweb

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (zeroweb *Zeroweb) reloadLogger() error {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = ""
	loglevelstr := zeroweb.Config.GetString("log.LogLevel")
	level, err := zerolog.ParseLevel(loglevelstr)
	if err != nil {
		log.Fatal().Err(err).Str("loglevel", loglevelstr).Msg("Bad Loglevel")
		return err
	}
	zerolog.SetGlobalLevel(level)
	zerolog.LevelFieldName = zeroweb.Config.GetString("log.LevelFieldName")
	zerolog.MessageFieldName = zeroweb.Config.GetString("log.MessageFieldName")
	zerolog.TimestampFieldName = zeroweb.Config.GetString("log.TimestampFieldName")
	return nil
}
