package zeroweb

import (
	"github.com/godofdream/fasthttp"
	"github.com/godofdream/jet"
	"github.com/google/go-cmp/cmp"
)

func (zeroweb *Zeroweb) reloadHTTP() error {
	zeroweb.View = jet.NewHTMLSet(zeroweb.Config.GetString("templates.folder"))

	// Routes
	zeroweb.Router.ServeFiles("/static/*filepath", zeroweb.Config.GetString("static.static_folder"))
	zeroweb.Router.ServeFiles("/css/*filepath", zeroweb.Config.GetString("static.css_folder"))
	zeroweb.Router.ServeFiles("/js/*filepath", zeroweb.Config.GetString("static.js_folder"))
	zeroweb.Router.ServeFiles("/fonts/*filepath", zeroweb.Config.GetString("static.fonts_folder"))

	// httpserver

	var httpconfig *fasthttp.Server
	var oldhttpconfig *fasthttp.Server = zeroweb.Server
	err := zeroweb.Config.UnmarshalKey("http", httpconfig)
	if err != nil {
		zeroweb.Log.Error().Err(err).Msg("unable to decode http config into struct")
		return err
	}

	if oldhttpconfig != nil && cmp.Equal(httpconfig, oldhttpconfig) {
		return nil // config didn't change for Webserver
	}

	zeroweb.Server = &fasthttp.Server{
		Handler: zeroweb.Router.Handler,
		Logger:  zeroweb.Log,
	}

	return nil
}
