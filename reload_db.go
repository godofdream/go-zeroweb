package zeroweb

import (
	"github.com/go-pg/pg"
	"github.com/rs/zerolog/log"
)

func (a *Zeroweb) reloadDB() error {
	if !a.Config.GetBool("db.enabled") {
		return nil
	}

	var config pg.Options
	var oldconfig *pg.Options = nil
	if a.DB != nil {
		oldconfig = a.DB.Options()
	}

	err := a.Config.Unmarshal(&config)
	if err != nil {
		a.Log.Fatal().Err(err).Msg("unable to decode into struct")
	}

	if oldconfig != nil &&
		config.Addr == oldconfig.Addr &&
		config.User == oldconfig.User &&
		config.Password == oldconfig.Password &&
		config.Database == oldconfig.Database &&
		config.MaxRetries == oldconfig.MaxRetries &&
		config.RetryStatementTimeout == oldconfig.RetryStatementTimeout &&
		config.MinRetryBackoff == oldconfig.MinRetryBackoff &&
		config.MaxRetryBackoff == oldconfig.MaxRetryBackoff &&
		config.DialTimeout == oldconfig.DialTimeout &&
		config.ReadTimeout == oldconfig.ReadTimeout &&
		config.WriteTimeout == oldconfig.WriteTimeout &&
		config.PoolSize == oldconfig.PoolSize &&
		config.PoolTimeout == oldconfig.PoolTimeout &&
		config.IdleTimeout == oldconfig.IdleTimeout &&
		config.MaxAge == oldconfig.MaxAge &&
		config.IdleCheckFrequency == oldconfig.IdleCheckFrequency {
		return nil // config didn't change for DB
	}

	//config.DialFunc //TODO cache dns, reuseport
	db := pg.Connect(&config)
	if db == nil {
		if a.DB == nil {
			log.Fatal().Err(err).Interface("config", config).Msg("connection to DB failed")
		}
		log.Error().Err(err).Interface("config", config).Msg("connection to DB failed (keeping old db connection)")
		return err
	}

	a.DB = db
	return nil
}
