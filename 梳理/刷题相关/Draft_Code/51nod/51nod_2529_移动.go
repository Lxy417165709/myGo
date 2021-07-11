package _1nod

import "fmt"

func main() {
	var n,k,s int
	fmt.Scan(&n,&k,&s)
	var l,h int
	h = (n-1)*k
	l = k

	if s>=l && s<=h{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}
