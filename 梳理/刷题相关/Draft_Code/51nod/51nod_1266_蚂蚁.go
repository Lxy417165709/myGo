

package _1nod

import (
	"fmt"
)

func getMax(a int,b int) int{
	if a>b{
		return a
	}
	return b
}
func getMin(a int,b int) int{
	if a>b{
		return b
	}
	return a
}

func main() {
	var L int
	var n int
	fmt.Scan(&n, &L)
	var x int
	m1,m2 := 0,0

	for i:=0;i<n ;i++  {
		fmt.Scan(&x)
		m1 = getMax(m1,getMax(x,L-x))
		m2 = getMax(m2,getMin(x,L-x))
	}
	fmt.Println(m2,m1)

}
