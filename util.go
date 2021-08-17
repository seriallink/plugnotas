package plugnotas

func In(value string, list ...string) bool {

	for _, element := range list {
		if value == element {
			return true
		}
	}

	return false
}

func NotIn(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return false
		}
	}

	return true

}

func StringPtr(value string) *string {
	return &value
}

func IntPtr(value int) *int {
	return &value
}

func FloatPtr(value float64) *float64 {
	return &value
}

func BoolPtr(value bool) *bool {
	return &value
}
