package _1nod

import "fmt"

var m [200050]int
var num [200050]int

func main() {
	n, k, a, g := 0, 0, 0, 0
	fmt.Scan(&n, &k, &a)

	fmt.Scan(&g)
	for i := 1; i <= g; i++ {
		fmt.Scan(&num[i])
	}
	l, r := 0, 0
	gs := (n+1) / (a+1)

	m[0] = 1
	m[n+1] = 1
	for i := 1; i <= g; i++ {
		m[num[i]] = 1
		for l = num[i] - 1; m[l] == 0; l-- {
		}
		for r = num[i] + 1; m[r] == 0; r++ {
		}
		gs = gs - ((r-l)/(a+1) - (num[i]-l)/(a+1) - (r-num[i])/(a+1))
		//fmt.Println(gs)
		if gs < k {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("-1")

}
/*
	总结
	1.  这题需要用到逆向思维
	2.  注意题意: 输入的n,k,a保证是能够在1×n的表格中放入k只大小为a的战舰，
		并且他们之间不重叠也不接触。不接触!!!
*/