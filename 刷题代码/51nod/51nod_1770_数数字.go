package _1nod

import "fmt"

func main() {
	//for i:=1;i<=9 ;i++  {
	//	for t:=1;t<=9 ;t++  {
	//		num := 0
	//		for k:=1;k<=4 ;k++  {
	//			num = num * 10 + i
	//		}
	//		num = num * t
	//		fmt.Printf("%d\t",num)
	//	}
	//	fmt.Println()
	//}
	T,a,b,d,n := 0,0,0,0,0
	fmt.Scan(&T)
	for ;T>0;T--{
		fmt.Scan(&a,&b,&d,&n)
		m := make(map[int]int)
		ans := a*b
		if ans<=9{
			m[ans] += n
		}else{
			h := ans / 10
			l := ans % 10
			if h+l <=9{
				m[h]++
				m[l]++
				m[h+l] = m[h+l] + n-1

			}else{
				if n==1{
					m[h]++
					m[l]++
				}else{
					m[h+1]++
					m[l]++
					m[(h+l)%10]++
					m[(h+l+1)%10] = m[(h+l+1)%10] + n - 2
				}



			}

		}
		fmt.Println(m[d])
	}


}
// 5 * 6
// 6 * 8
// 7 * 4
// 7 * 7
// 7 * 8
// 8 * 8

/*
	总结
	1. 这是一道找规律的题目~...

*/