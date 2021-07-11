package main

import "fmt"

func minimumSwap(s1 string, s2 string) int {
	l:=0
	r:=len(s1)-1
	ans:=0
	for l<r{
		for l<r  {
			if s1[l]!=s2[l]{
				break
			}
			l++
		}
		for r>l  {
			if s1[r]!=s2[r]{
				break
			}
			r--
		}

		if l==r{
			if s1[l]!=s2[l]{
				return -1
			}else{
				break
			}
		}else{
			if l<r{
				if s1[l]==s1[r]{
					ans+=1
				}else{
					ans+=2
				}
			}
		}
		l++
		r--
	}
	return ans



}
func main() {
	fmt.Println(minimumSwap("xxyyxyxyxx","xyyxyxxxyx"))
}
