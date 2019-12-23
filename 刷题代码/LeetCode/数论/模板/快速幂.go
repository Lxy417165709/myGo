package main

import "fmt"

func fastPow(a,n,m int) int{
	sum:=1
	for ;n!=0 ;  {
		a %=m
		if n&1==1 {
			sum=sum*a%m
		}
		a = a * a
		n>>=1
	}
	return sum
}

func main() {
	fmt.Println(fastPow(787,5,10057867))
}
