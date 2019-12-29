package main

/*
	现在假设有personCount人(标号[0,personCount-1])，他们每个人最多拥有30类物品(标号[0,29])，
	每类物品数量范围为[0,10^6]。
	现在给出他们拥有物品的情况，每一种情况用一个长度为3的切片表示。
	如 [0,5,10]，表示第0个人，拥有第5类物品10个。
	如 [5,10,1000]，表示第5个人，拥有第10类物品1000个。
	现在给出x,y，我想知道第x个人和第y个人 他们有多少类相同的、多少类不同的、以及他们一共有多少类物品。

	以上输入都在合法范围内。
*/

// 解题代码
func solve(personCount int,situations [][]int,x,y int) (sameKinds,unsameKinds,allKinds int){
	personOwn := make([]int,personCount)	// 用于表示每个人的物品拥有情况(不关心拥有多少个，0个表示没有，否则表示有)
	// 		这里使用32位整数表示人们拥有物品的情况:
	// 		假如 personOwn[i] = 7,  这表示第i个人拥有第0、第1、第2类物品。
	// 		解释:	对于32位整数来说，7的二进制是:  00000000000000000000000000000111，从右到左算，标1的位置表示拥有该类物品，否则为不拥有。
	// 		假如 personOwn[i] = 9,  这表示第i个人拥有第0、第3类物品。
	// 		解释:	对于32位整数来说，9的二进制是:  00000000000000000000000000001001，从右到左算，标1的位置表示拥有该类物品，否则为不拥有。

	for i:=0;i<len(situations);i++{
		if situations[i][2]!=0{
			personOwn[situations[i][0]] &= 1<<uint8(situations[i][1])
		}
	}
	// 相同的物品表示2个人该二进制位上都是1，所以使用 & 运算
	sameKinds = countOne(personOwn[x]&personOwn[y])

	// 不同的物品表示2个人该二进制位上是不同的，所以使用 ^ 运算
	unsameKinds = countOne(personOwn[x]^personOwn[y])

	// 对于一共有多少类物品，两个人在该位上有一个1就可以了。
	allKinds = countOne(personOwn[x]|personOwn[y])

	// 例子:
	// 如 				personOwn[x] == 00000000000000000000000000000111
	//	  				personOwn[y] == 00000000000000000000000000001001

	// 则  personOwn[x]&personOwn[y] == 00000000000000000000000000000001	  	表示x,y共同拥有第0类物品，所以只有1类相同物品
	//     personOwn[x]^personOwn[y] == 00000000000000000000000000001110		表示x,y第1、2、3类物品拥有情况不同，所有有3类不同物品
	//     personOwn[x]|personOwn[y] == 00000000000000000000000000001111		表示x,y一共拥有第0、1、2、3物品，所以一共有4类物品。

	return
}

// 计算a有多少个1
func countOne(a int) int{
	count := 0
	for a != 0{
		count++
		a = a &(a-1)
	}
	return count
}


