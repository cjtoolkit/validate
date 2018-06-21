package vFloat

func toBoolMap(values []float64) map[float64]bool {
	m := map[float64]bool{}
	for _, value := range values {
		m[value] = true
	}
	return m
}
