package main

func isIsomorphic(s string, t string) bool {
	return judge(s,t) && judge(t,s)
}

func judge(s string,t string) bool{
	mp:=make(map[uint8]uint8)
	for i:=0;i<len(s);i++{
		if x,ok:=mp[s[i]];ok{
			if x != t[i]{
				return false
			}
		}else{
			mp[s[i]]=t[i]
		}
	}
	return true
}
/*
	总结
	1. 该题主要就是要求，是否2个字符串是唯一映射的，即每个s中的每一个字符唯一对应t中的每一个，t中的每一个字符也唯一对应s中的每一个。
	2. 我是使用map解决这个题目的..
	3. 还有一种写法就是先获取s->t的map，然后获取t->s的map，之后判断是否 mp[t]->s,mp[s]->t (这里指的都是字符)
		全部满足则为true,否则为false
*/
