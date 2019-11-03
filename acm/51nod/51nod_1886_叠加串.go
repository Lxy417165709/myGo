package _1nod

import (
	"fmt"
)
const bignum = 1152921504606846975

func solov(number int64,a int64,lay int64) int64{
	//fmt.Println(number,"\t",a,"\t",lay)
	mid := number / 2 + 1

	if a == mid{
		if lay%2 == 1{
			return 1
		}else{
			return 0
		}
	}

	//if number <= 3{
	//	if a==3{
	//		if lay%2 == 1{
	//			return 0
	//		}else{
	//			return 1
	//		}
	//	}
	//	if a==1{
	//		if lay%2 == 1{
	//			return 1
	//		}else{
	//			return 0
	//		}
	//	}
	//
	//}

	if mid > a{
		return solov(mid-1,a,lay)
	}else{
		//return solov(mid-1,mid - (a-mid),lay+1)
		return solov(mid-1,mid - int64(a - mid),lay+1)

		//solov(mid-1,mid - int64(float64(a - mid)),lay+1) 这样是错误的!...
		// 注意浮点数误差..浮点误差会使这题错误....哭了
	}




}




func fun (){
	S:="0"
	//fmt.Println(S)
	//var l int64=1
	//var i int64=1
	//for i=1;l<1E18;i++  {
	//	l = l*2 + 1
	//
	//}
	//println(l,i)

	for i:=0;i<3;i++  {
		temp:=change(S)

		S = S+"0" + temp
		fmt.Println(S)
	}
	fmt.Println(S)

	//num:=0

	//for i:=0;i<200 ;i++  {
	//	if S[i]=='1'{
	//		fmt.Print(i,"\t")
	//		num++
	//		if num%5==0{
	//			fmt.Println()
	//		}
	//	}
	//
	//}

}
func change(s string)string  {
	b := []byte(s)

	for i:=0;i<len(b);i++  {
		if b[i]=='0'{
			b[i] = '1'
		}else{
			b[i]='0'
		}
	}
	i:=0
	j:=len(b)-1
	for ;i<=j ;  {
		b[i],b[j] = b[j],b[i]
		i++
		j--
	}
	return string(b)

}
func main() {
	//fun ()
	var n,T int64 =0,0


	fmt.Scan(&T)
	for ;T>0 ;T--  {
		fmt.Scan(&n)
		fmt.Println(solov(bignum,n,0))
	}



}
/*
	总结
	1. 注意浮点数误差..浮点误差会使这题错误....哭了
	2. 这题目可以先模拟一下找字符串规律，然后用二分递归思维解答
	3. bignum = 1152921504606846975 超过1E18了!

*/