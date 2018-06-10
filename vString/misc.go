package vString

func toBoolMap(values []string) map[string]bool {
	m := map[string]bool{}
	for _, value := range values {
		m[value] = true
	}
	return m
}
