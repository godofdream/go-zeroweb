package zeroweb

func (a *Zeroweb) Reload() error {
	//TODO reload on config change
	if err := a.ReloadLogger(); err != nil {
		return err
	}
	if err := a.ReloadDB(); err != nil {
		return err
	}
	if err := a.ReloadHTTP(); err != nil {
		return err
	}
	return nil
}
