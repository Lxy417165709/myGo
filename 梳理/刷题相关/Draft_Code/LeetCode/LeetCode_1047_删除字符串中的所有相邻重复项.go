package 双指针

func removeDuplicates(S string) string {
	l:=-1
	r:=0
	ans:=""
	for r<len(S){
		if l!=-1 &&  ans[l]==S[r]{
			ans = ans[:l]
			l--
		}else{
			ans+=string(S[r])
			l++
		}
		r++
	}
	return ans
}


/*
	总结
	1. 第一次做的时候以为只要重复了就全部删除了，然后WA后才知道了题意....所以要细心呀
	2. 第二次看题意后就使用双指针AC了，如下：
			我维护2个指针,l,r，l指向答案的第一个字符(如果答案是空，则l==-1)
		    之后r向前推进，每次都和l指向的字符做对比，对比之下进行相应操作，最后就
			可以得出答案了！
	3. 答案的数据结构是栈 0.0..
*/
