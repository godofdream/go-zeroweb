package zeroweb

import (
	"github.com/godofdream/fasthttp"
	"github.com/godofdream/jet"
	"github.com/google/go-cmp/cmp"
)

func (zeroweb *Zeroweb) reloadHTTP() error {
	// Jet templating
	zeroweb.View = jet.NewHTMLSet(zeroweb.Config.GetString("templates.Folder"))

	// Routes
	zeroweb.Router.ServeFiles("/static/*filepath", zeroweb.Config.GetString("static.StaticFolder"))
	zeroweb.Router.ServeFiles("/css/*filepath", zeroweb.Config.GetString("static.CSSFolder"))
	zeroweb.Router.ServeFiles("/js/*filepath", zeroweb.Config.GetString("static.JSFolder"))
	zeroweb.Router.ServeFiles("/fonts/*filepath", zeroweb.Config.GetString("static.FontsFolder"))

	// httpserver

	var httpconfig fasthttp.Server
	oldhttpconfig := zeroweb.Server
	err := zeroweb.Config.UnmarshalKey("http", &httpconfig)
	if err != nil {
		zeroweb.Log.Error().Err(err).Msg("unable to decode http config into struct")
		return err
	}

	if oldhttpconfig != nil && cmp.Equal(&httpconfig, oldhttpconfig) {
		return nil // config didn't change for Webserver
	}

	httpconfig.Handler = zeroweb.Router.Handler
	httpconfig.Logger = zeroweb.Log
	zeroweb.Server = &httpconfig

	return nil
}
