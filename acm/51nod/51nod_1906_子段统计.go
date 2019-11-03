package _1nod

import "fmt"
const mod = 1610612741

func main() {
	m := map[int] int{}
	l,k:=0,0
	fmt.Scan(&l,&k)
	s:=""
	fmt.Scan(&s)
	cnt:=0
	for i:=0;i<=len(s)-l ;i++  {
		var num = 0
		for t:=0;t<l ;t++  {
			num = (num * 30 + int(s[i+t]) - 96)%mod
		}
		_,ok := m[num]
		if ok == false{
			m[num] ++
			cnt++
		}

	}
	fmt.Println(cnt)
	//fmt.Println(m)
}
/*
	总结
	1. 注意Hash映射的时候，可能出现前导0,比如 baaaa,b,如果a表示为0，那么他们的哈希值是一样的
		(在这里是左低位，右高位)

*/