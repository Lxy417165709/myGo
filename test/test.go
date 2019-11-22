package main

import "fmt"
// total小于str长度时，我返回的是截取后的str中，第一个字符出现的次数
func repeatedString(str string,total int) int{


	if str==""{
		return 0
	}
	char:=str[0]

	times:=0
	for i:=0;i<len(str);i++{
		if str[i]==char{
			times++
		}
	}
	ans:= times*(total/len(str))
	finallyIndex:=total%len(str)

	for i:=0;i<finallyIndex;i++{
		if str[i]==char{
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(repeatedString("aba",2))
}
