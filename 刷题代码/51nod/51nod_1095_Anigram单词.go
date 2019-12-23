package _1nod

import (
	"fmt"
	"sort"
)

type rSlice []rune

func (r rSlice) Len() int {
	return len(r)
}
func (r rSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r rSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func main() {
	var N, Q int
	var s string
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	fmt.Scan(&N)
	for ; N > 0; N-- {
		fmt.Scan(&s)
		m1[s]++
		r := []rune(s[:])
		sort.Sort(rSlice(r))
		s2 := string(r)
		m2[s2]++
	}
	fmt.Scan(&Q)
	for ; Q > 0; Q-- {
		fmt.Scan(&s)
		r := []rune(s[:])
		sort.Sort(rSlice(r))
		s2 := string(r)
		fmt.Println(m2[s2] - m1[s])
	}
}
