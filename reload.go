package zeroweb

func (a *Zeroweb) Reload() error {
	//TODO reload on config change
	if err := a.reloadLogger(); err != nil {
		return err
	}
	if err := a.reloadDB(); err != nil {
		return err
	}
	if err := a.reloadHTTP(); err != nil {
		return err
	}
	return nil
}
