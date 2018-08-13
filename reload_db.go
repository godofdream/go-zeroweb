package zeroweb

import (
	"github.com/go-pg/pg"
	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog/log"
)

func (zeroweb *Zeroweb) reloadDB() error {
	if !zeroweb.Config.GetBool("db.Enabled") {
		return nil
	}

	var config pg.Options
	var oldconfig *pg.Options = nil
	if zeroweb.DB != nil {
		oldconfig = zeroweb.DB.Options()
	}

	err := zeroweb.Config.UnmarshalKey("db", &config)
	if err != nil {
		zeroweb.Log.Error().Err(err).Msg("unable to decode db config into struct")
		return err
	}

	if oldconfig != nil && cmp.Equal(config, oldconfig) {
		return nil // config didn't change for DB
	}

	//config.DialFunc //TODO cache dns, reuseport
	db := pg.Connect(&config)
	if db == nil {
		if zeroweb.DB == nil {
			log.Fatal().Err(err).Interface("config", config).Msg("connection to DB failed")
		}
		log.Error().Err(err).Interface("config", config).Msg("connection to DB failed (keeping old db connection)")
		return err
	}

	zeroweb.DB = db
	return nil
}
