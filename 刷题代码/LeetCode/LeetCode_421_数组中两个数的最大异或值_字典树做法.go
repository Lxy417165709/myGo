package main

import "strconv"

func findMaximumXOR(nums []int) int {
	root:=&DicTree{-1,[2]*DicTree{}}

	for i:=0;i<len(nums);i++{
		root.insert(nums[i])
	}
	res:=0
	for i:=0;i<len(nums);i++{
		res = max(res,root.getMaximumXOR(nums[i]))
	}
	return res

}



func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

type DicTree struct{
	val int
	son [2]*DicTree
}

func (this *DicTree)insert(number int){
	str:=""
	for i:=31;i>=0;i--{

		director :=number & (1<<uint8(i))   // 注意优先级啊

		if director!=0{
			director = 1
		}
		str+=strconv.Itoa(director)
		if this.son[director] == nil{
			this.son[director] = &DicTree{director,[2]*DicTree{}}
		}
		this = this.son[director]
	}

	this.son[0] = &DicTree{number,[2]*DicTree{}}
}

func (this *DicTree)getMaximumXOR(number int) int{
	for t:=31;t>=0;t--{

		director :=number & (1<<uint8(t))   // 注意优先级啊
		if director!=0{
			director = 1
		}
		if this.son[director^1]!=nil{
			this = this.son[director^1]
		}else{
			this = this.son[director]
		}
	}
	return number^this.son[0].val
}


func main() {

}

/*
	总结
	1. 第一次这题我是看官方题解，然后采用贪心的做法AC的
	2. 之后看了字典树的，然后理解之后自己也动手写
	3. 代码完成之后，发现有BUG，找了好久之后才发现是运算优先级的问题
		 number & (1<<uint8(i))    &的优先级比<<高，所以要加括号。
		 找这个BUG找了好久 0.0..
*/