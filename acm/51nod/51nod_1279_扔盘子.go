package _1nod

import "fmt"

var (
	w [50500]int
)
func main() {
	var m,n int
	fmt.Scan(&n,&m)
	fmt.Scan(&w[0])
	for i:=1;i<n ;i++  {
		fmt.Scan(&w[i])
		if w[i-1]<w[i]{
			w[i] = w[i-1]
		}
	}
	var x int
	ans:=0
	for i:=0;i<m ;i++  {
		fmt.Scan(&x)
		n--
		for ; n>=0 && w[n]<x ;  {
			n--
		}
		if n>=0 {
			ans++
		}
	}
	fmt.Println(ans)
}
/*
	总结
	1. 这是一道思维题
	2. 所用知识点是: 模型变换和逆向思维~

*/