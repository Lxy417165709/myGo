package main

import "sort"



func nthUglyNumber(n int, a int, b int, c int) int {
	tmp := []int{a, b, c}
	sort.Ints(tmp)

	l := tmp[0]              // 下界
	r := 2000000000000000000 // 上界
	//ans := r
	// 可AC二分，但是我看了别人后，发现他们更优美，如下
	//for l <= r {
	//	mid := (l + r) / 2
	//	cnt := count(a, b, c, mid)
	//	if cnt == n {
	//		// 这部分单独也可以
	//		// if mid%a!=0 &&mid%b!=0 && mid%c!=0{
	//		//     r = mid - 1
	//		// }else{
	//		//     ans = min(ans,mid)
	//		//     r = mid - 1
	//		// }
	//
	//		// 这部分单独也可以
	//		ans = min(ans, mid) // 要注意，可能有多个值对应同一个位置
	//		r = mid - 1
	//
	//	} else {
	//		if cnt > n {
	//			r = mid - 1
	//		} else {
	//			l = mid + 1
	//		}
	//	}
	//}

	// 返回最小满足条件的数
	for l<r{
		mid := l+(r-l)/2
		cnt := count(a,b,c,mid)
		if cnt>=n{
			r = mid
		}else{
			l = mid + 1
		}
	}
	return l


}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 必须保证arr长度大于等于1
// 第一次WA是因为lcm写错了。。
func lcm(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	x := lcm(arr[:len(arr)-1]...)
	GCD := gcd(x, arr[len(arr)-1])
	return x / GCD * arr[len(arr)-1]
}

// 返回它是第多少个丑数（它可能不是丑数，此时返回的是比它小且最靠近它的丑数的号码(第多少个)）
func count(a int, b int, c int, num int) int {
	// 这里还可以优化，cnt_x 都是可以固定得到的值。。
	cnt_a, cnt_b, cnt_c := num/lcm(a), num/lcm(b), num/lcm(c)
	cnt_ab, cnt_ac, cnt_bc := num/lcm(a, b), num/lcm(a, c), num/lcm(b, c)
	cnt_abc := num / lcm(a, b, c)

	return cnt_a + cnt_b + cnt_c - cnt_ab - cnt_ac - cnt_bc + cnt_abc
}



// 这个也可以AC,是我自己对自己的二分的优美化
//func nthUglyNumber(n int, a int, b int, c int) int {
//	tmp := []int{a, b, c}
//	sort.Ints(tmp)
//
//	l := tmp[0]              // 下界
//	r := 2000000000000000000 // 上界
//	for l <= r {
//		mid := (l + r) / 2
//		cnt := count(a, b, c, mid)
//		if cnt == n {
//			r = mid - 1
//		} else {
//			if cnt > n {
//				r = mid - 1
//			} else {
//				l = mid + 1
//			}
//		}
//	}
//	return l
//
//}



func main() {

}

/*
	总结
	1. 这题如果按照动态规划做的话肯定超时
	2. 要正确做出这题需要用到数论知识。。
	3. 这思维太棒了，不是从n得到数，而是通过假设数获取他是第x个丑数，之后再通过x与n的关系变化进行二分查找数。
*/
