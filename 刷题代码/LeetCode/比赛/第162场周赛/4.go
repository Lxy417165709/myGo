package main
func maxScoreWords(words []string, letters []byte, score []int) int {
	mp:=make(map[byte]int)
	// 获得字符资源
	for i:=0;i<len(letters);i++{
		mp[letters[i]]++
	}
	return solve(words,mp,score)
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

// 深复制一个map ...
func newMap(mp map[byte]int) map[byte]int{
	tmp:=make(map[byte]int)
	for k,v := range mp{
		tmp[k] = v
	}
	return tmp
}

// 对于一个单词，我们可以选择不构成它or构成它,
// 不构成它: solve(words[1:],nowMp,score)  此时不费字符资源
// 构成它  : getSum(words[0],nowMp,score) + solve(words[1:],nowMp,score) 此时要费字符资源
func solve(words []string,mp map[byte]int,score []int) int{
	if len(words)==0{
		return 0
	}
	nowMp := newMap(mp)

	// 这会存在重叠子问题（当当前的字符资源无法构成单词时）
	return max(
		solve(words[1:],nowMp,score),
		getSum(words[0],nowMp,score) + solve(words[1:],nowMp,score),    // getSum会改变nowMap
	)
}

// 获取target的价值
func getSum(target string,mp map[byte]int,score []int) int{
	ans:=0

	for i:=0;i<len(target);i++{
		if mp[target[i]]!=0{
			mp[target[i]]--
			ans+=score[int(target[i]-'a')]
		}else{
			return 0   // 表示当前字符无法构成该单词，返回无穷小,0也可以的
		}
	}

	return ans
}

func main() {

}

/*
	总结
	1. 要注意map的深复制和浅复制问题..不然会出现很难察觉的错...
*/