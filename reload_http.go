package zeroweb

import "github.com/godofdream/jet"

func (a *Zeroweb) ReloadHTTP() error {
	a.View = jet.NewHTMLSet("./templates")
	return nil
}
