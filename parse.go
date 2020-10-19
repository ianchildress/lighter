package lighter

func Parse() error {
	mutex.Lock()
	defer mutex.Unlock()

	for _, f := range flags {
		if err := f.parse(); err != nil {
			return err
		}
	}

	return nil
}
