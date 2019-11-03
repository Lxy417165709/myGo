package main

import "fmt"

func isPrime(n int) bool{
	if n==1 || n==0 {
		return false
	}
	for i:=2;i*i<=n ;i++  {
		if n%i==0 {
			return false
		}
	}
	return true
}

func jc(n int,m int) int{
	if n==0 || n==1 {
		return 1
	}
	return n*jc(n-1,m)%m
}

func numPrimeArrangements(n int) int {
	x:=0
	for i:=1;i<=n ;i++  {
		if isPrime(i)==true {
			x++
		}
	}

	m:=1000000007
	ans:=jc(x,m)*jc(n-x,m)%m
	return ans
}
func main() {
	fmt.Println(numPrimeArrangements(4))
}
/*
	总结
	1. 这是一道简单题，考查排列组合0.0..
*/