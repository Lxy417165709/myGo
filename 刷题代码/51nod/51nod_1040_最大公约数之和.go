package _1nod

import "fmt"

func euler(n int) int{
	ans := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			ans = ans / i * (i - 1)
			for ; n%i == 0; {
				n = n / i
			}
		}
	}
	if n > 1 {
		ans = ans / n * (n - 1)
	}
	return ans

}

func main() {
	n:=0
	fmt.Scan(&n)
	ans:=0
	for i:=1;i*i<=n;i++{
		if n%i == 0{
			if i*i == n {
				ans += i*euler(n/i)
			}else{
				ans += i*euler(n/i) + (n/i)*euler(i)
			}

		}

	}

	fmt.Println(ans)

}

/*
	总结
	1.  设Z(x,y)为 [1,x]中与x最大公约数为y的函数
		则Z(x,y) == Z(x/y,1) == euler(x/y)
	2.  求取因子，假如x有因子y,那么x/y也是x的因子，所以用sqrt(n)的复杂度就可以获取它的
		所有因子。

*/