package zeroweb

func (a *Zeroweb) Reload() {
	//TODO reload on config change
	a.reloadLogger()
	a.reloadDB()
	a.reloadHTTP()
}
