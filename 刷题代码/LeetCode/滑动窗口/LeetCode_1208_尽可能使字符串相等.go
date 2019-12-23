package main


func equalSubstring(s string, t string, maxCost int) int {
	l,r:=0,0
	ans:=0
	for l<=r && r<=len(s)-1 {

		except:=abs(int(s[r])-int(t[r]))
		maxCost-=int(except)
		if maxCost >= 0{
			ans = max(ans,r-l+1)
		}else{
			for l<=r && maxCost<0{
				maxCost+=abs(int(s[l])-int(t[l]))
				l++
			}
		}
		r++
		/*
			except:=abs(int(s[r])-int(t[r]))
			maxCost-=int(except)

			for l<=r && maxCost<0{
				maxCost+=abs(int(s[l])-int(t[l]))
				l++
			}
			ans = max(ans,r-l+1)
			r++

			这样写更优美

		*/
	}
	return ans
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func abs(a int) int{
	if a>0{
		return int(a)
	}
	return int(-a)
}

func main() {

}
/*
	总结
	1. 这是一道滑动窗口题
	2. 我自认为这个滑动窗口优美了一些 0.0
	3. 我不是计算滑动窗口之中的代价的，我是假如装下滑动窗口的字符，maxCost还剩下多少，
	   如果是>=0，则继续推进，否则就要从左边开始缩了 0.0
	4. 其实左边缩的过程可以采用二分优化 0.0

*/