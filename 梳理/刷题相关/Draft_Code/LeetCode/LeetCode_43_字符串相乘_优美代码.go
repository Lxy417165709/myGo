package main

func multiply(num1 string, num2 string) string {
	len1,len2:=len(num1),len(num2)
	result := make([]int,len1+len2)

	for i:=len1-1;i>=0;i--{
		for t:=len2-1;t>=0;t--{
			sum:=int((num1[i]-'0')*(num2[t]-'0'))
			tmp:=sum+result[i+t+1]
			result[i+t+1]=tmp%10
			result[i+t] +=tmp/10
		}
	}

	// 接下来是为了去除前导0
	begin:=0
	for begin<len(result) && result[begin]==0{
		begin++
	}
	return intSliceToString(result[begin:])
}

func intSliceToString(slice []int) string{
	if len(slice)==0{
		return "0"
	}
	result:=""
	for i:=0;i<len(slice);i++{
		result+=string(slice[i]+'0')
	}
	return result
}
