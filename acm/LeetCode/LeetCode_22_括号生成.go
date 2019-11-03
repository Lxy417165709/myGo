package main

import "fmt"
var ans map[int] []string
var m map[string] bool
func generateParenthesis(n int) []string {
	ans=make(map[int] []string)
	m = make(map[string] bool)
	ans[1] = []string{"()"}
	return solve(n)
}

func solve(n int) []string{

	for i:=2;i<=n ;i++  {
		temp :=make([]string, len(ans[i-1]))
		copy(temp,ans[i-1])

		for t:=0;t< len(temp);t++  {
			s :=temp[t]
			s = "()" + s
			s_arr := []byte(s)
			index:=1

			if m[s]==false {
				m[s]=true
				ans[i] = append(ans[i], s)
			}

			for u:=2;u< len(s_arr);u++  {
				if s_arr[u]=='(' {
					s_arr[u],s_arr[index] = s_arr[index],s_arr[u]
					index = u
					u = index + 1
					if m[string(s_arr)]==false {
						m[string(s_arr)]=true
						ans[i] = append(ans[i],string(s_arr))
					}
				}
			}
		}
	}
	return ans[n]
}



//func judge(s string) bool{
//
//	for i:=1;i< len(s);i = i + 2  {
//		if s[i-1:i+1]!="()" {
//			return true
//		}
//	}
//	return false
//
//}

func test( slice []string) (int,int){
	for i:=0;i< len(slice);i++  {
		for t:=i+1;t< len(slice);t++  {
			if slice[i]==slice[t] {
				return i,t
			}
		}
	}
	return -1,-1
}
// 递归的(错误)
//func solve(n int) []string{
//	slice := make([]string,0)
//
//	if n==1 {
//		return []string{"()"}
//	}else{
//		slice1 := solve(n-1)
//		for i:=0;i< len(slice1);i++  {
//			if judge(slice1[i]) == true {
//				slice = append(slice, slice1[i] + "()")
//			}
//			slice = append(slice, "()" + slice1[i])
//			slice = append(slice,"(" + slice1[i] + ")")
//		}
//		return slice
//	}
//}
func main() {
	fmt.Println(test(generateParenthesis(4)))
	fmt.Println(generateParenthesis(4))
}

/*
	总结
	1. 该问题不能使用我自己写的递归...(())()3对转4对时 可以是 (())(())
	2. 使用map AC了这题，但是太费空间了
	3. 官方有更好的题解，他用的是回溯递归
*/