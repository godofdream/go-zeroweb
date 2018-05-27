package zeroweb

import "github.com/godofdream/jet"

func (a *Zeroweb) ReloadHTTP() {
	a.View = jet.NewHTMLSet("./templates")
}
