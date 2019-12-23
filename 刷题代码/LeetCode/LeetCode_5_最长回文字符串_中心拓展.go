package main


func longestPalindrome(s string) string {
	if len(s)<=1{
		return s
	}

	begin,end:=0,0
	for i:=0;i<len(s);i++{
		len1:=expand(s,i,i)
		len2:=expand(s,i,i+1)
		length:=max(len1,len2)

		// end-begin+1就是此时回文字符串的长度
		// length>end-begin+1 表示要大于此时的回文字符串长度才更新
		// length>end-begin 等于 length>=end-begin+1   表示≥此时的回文字符串长度都更新
		if length>end-begin+1{
			begin=i-(length-1)/2
			end=i+length/2
		}
	}
	return s[begin:end+1]

}


func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func expand(s string,left int,right int) int{

	for left>=0 && right<len(s) && s[left]==s[right]{
		left--
		right++
	}

	return right-left-1 // 注意是 -1，因为回文字符串此时长度是 right-1 - (left+1) + 1
}



/*
	总结
	1. 注意中心拓展的细节 0.0..找了好久
*/