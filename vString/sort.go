package vString

import "sort"

func searchStrings(l int, a []string, x string) int {
	return sort.Search(l, func(i int) bool { return a[i] >= x })
}
