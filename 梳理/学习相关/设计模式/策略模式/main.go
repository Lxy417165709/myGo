package main

import "pattern"

func main(){
	person := pattern.Person{}

	person.EatRice()

	person.SetEatWay(pattern.Strategy1{})
	person.EatRice()

	person.SetEatWay(pattern.Strategy2{})
	person.EatRice()

	person.SetEatWay(pattern.Strategy3{})
	person.EatRice()

	// 吃饭是什么技能？我不会阿
	// 站着吃饭！
	// 坐着吃饭！
	// 倒立着吃饭！
}

// 做一件事情可以有很多的方法，这些方法就是策略，而策略模式就是把这些策略封装到一个类中、
