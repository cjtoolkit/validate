package vInt

type sortInt64 []int64

func (p sortInt64) Len() int           { return len(p) }
func (p sortInt64) Less(i, j int) bool { return p[i] < p[j] }
func (p sortInt64) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
