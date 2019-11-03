package _1nod

import "fmt"

var (
	one  [1000010]int
	zero [1000010]int
)

func main() {
	var s string

	fmt.Scan(&s)

	oneNumber := 0
	zeroNumber := 0
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '1' {
			oneNumber ++
			one[i] = one[i-1] + 1
			zero[i] = zero[i-1]
		} else {
			zeroNumber ++
			one[i] = one[i-1]
			zero[i] = zero[i-1] + 1
		}
	}
	//// 全是1和全是0时
	//if oneNumber == len(s) {
	//	fmt.Println(len(s))
	//	return
	//}
	//if zeroNumber == len(s) {
	//	fmt.Println(0)
	//	return
	//}

	// 统一起来了~
	for i := 0; i <= len(s); i++ {
		if oneNumber-one[i] == zero[i] {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)

}
