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


func main() {
	fmt.Println(isPrime(1000000007))
}
