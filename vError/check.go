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
