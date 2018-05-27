package zeroweb

func (a *Zeroweb) Reload() {
	//TODO reload on config change
	a.ReloadLogger()
	a.ReloadDB()
	a.ReloadHTTP()
}
