package main

import (
	"fmt"
	"sort"
)

func maxNumberOfBalloons(text string) int {

	if text=="" {
		return 0
	}

	m:=make(map[uint8]int)
	for i:=0;i< len(text);i++  {
		m[text[i]]++
	}

	slice:=make([]int,0)
	slice = append(slice, m['b'])
	slice = append(slice, m['a'])
	slice = append(slice, m['l']/2)
	slice = append(slice, m['o']/2)
	slice = append(slice, m['n'])
	sort.Ints(slice)
	return slice[0]

}
func main() {
	fmt.Println(maxNumberOfBalloons("leetcode"))
}
