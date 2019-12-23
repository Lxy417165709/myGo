package main


func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	count := make(map[string]int)
	ans := 0
	for i:=minSize-1;i<len(s);i++{
		if calculate(s[i+1-minSize:i+1])<=maxLetters{
			count[s[i+1-minSize:i+1]]++
			ans = max(ans,count[s[i+1-minSize:i+1]])
		}
	}
	return ans
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}


func calculate(s string) int{
	count := make(map[uint8]int)
	for i:=0;i<len(s);i++{
		count[s[i]]++
	}
	return len(count)
}

func main() {

}