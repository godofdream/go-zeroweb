package zeroweb

import "github.com/godofdream/jet"

func (a *Zeroweb) reloadHTTP() error {
	a.View = jet.NewHTMLSet("./templates")
	return nil
}
