package main

func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}

func gcd(a,b int) int{
	if b==0 {
		return a
	}
	return gcd(b,a%b)
}

func quickMod(a,n,m int) int{
	ans:=1
	for n!=0{
		a=a%m
		if n&1==1 {
			ans=ans*a%m
		}
		a = a*a
		n = n >> 1
	}
	return ans
}

//func main() {
//	fmt.Println(quickMod(10,0,7))
//}