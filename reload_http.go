package zeroweb

import "github.com/godofdream/jet"

func (a *Zeroweb) reloadHTTP() {
	a.View = jet.NewHTMLSet("./templates")
}
