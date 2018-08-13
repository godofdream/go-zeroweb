package zeroweb

func (zeroweb *Zeroweb) Reload() error {
	//TODO reload on config change
	if err := zeroweb.reloadLogger(); err != nil {
		return err
	}
	if err := zeroweb.reloadDB(); err != nil {
		return err
	}
	if err := zeroweb.reloadHTTP(); err != nil {
		return err
	}
	return nil
}
