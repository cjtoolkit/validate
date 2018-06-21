package vUint

func toBoolMap(values []uint64) map[uint64]bool {
	m := map[uint64]bool{}
	for _, value := range values {
		m[value] = true
	}
	return m
}
