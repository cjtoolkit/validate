package vInt

func toBoolMap(values []int64) map[int64]bool {
	m := map[int64]bool{}
	for _, value := range values {
		m[value] = true
	}
	return m
}
