package vError

func CheckErr(errs ...error) bool {
	for _, err := range errs {
		if nil != err {
			return false
		}
	}

	return true
}

func CheckBool(boolValues ...bool) bool {
	for _, boolValue := range boolValues {
		if !boolValue {
			return false
		}
	}

	return true
}

func Must(err error) {
	if nil != err {
		panic(err)
	}
}

func CleanError(err error) error {
	switch err := err.(type) {
	case Errors:
		dupeList := map[string]bool{}
		cleanErr := Errors{}

		for _, e := range err {
			if ve, ok := e.(ValidationError); ok {
				key := ve.Type + ve.Format
				if !dupeList[key] {
					dupeList[key] = true
					cleanErr = append(cleanErr, ve)
				}
			}
		}

		return cleanErr
	}

	return err
}

func MergeErrors(errs ...error) error {
	collector := NewErrorCollector()
	for _, err := range errs {
		collector.Collect(err)
	}

	return CleanError(collector.GetErrors())
}
