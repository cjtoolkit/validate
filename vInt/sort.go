package vInt

import "sort"

type sortInt64 []int64

func (p sortInt64) Len() int           { return len(p) }
func (p sortInt64) Less(i, j int) bool { return p[i] < p[j] }
func (p sortInt64) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func searchInt64s(l int, a []int64, x int64) int {
	return sort.Search(l, func(i int) bool { return a[i] >= x })
}
