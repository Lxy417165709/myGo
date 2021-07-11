package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n<0 {
		x = 1/x
		n = -n
	}
	return quickPow(x,n)
}

func quickPow(x float64,n int) float64{
	sum:=1.0

	for ;n!=0 ;  {
		if n%2==1 {
			sum=sum*x
		}
		x=x*x
		n = n /2
	}
	return sum
}
func main() {
	fmt.Println(myPow(2,10))
	fmt.Println(myPow(2.1,1))
	fmt.Println(myPow(1,-2))
}
