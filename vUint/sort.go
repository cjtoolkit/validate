package vUint

type sortUint64 []uint64

func (p sortUint64) Len() int           { return len(p) }
func (p sortUint64) Less(i, j int) bool { return p[i] < p[j] }
func (p sortUint64) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
