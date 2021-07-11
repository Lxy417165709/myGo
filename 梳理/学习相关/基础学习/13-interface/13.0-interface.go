package main

import "fmt"

type Animal interface{
	eat(string)
}
type Person struct {
}
func (person Person)eat(food string){
	fmt.Println("人吃"+food)
}

type Dog struct {
}
func (dog Dog)eat(food string){
	fmt.Println("狗舔"+food)
}

func main(){
	var animal Animal
	animal = new(Person)
	animal.eat("肉肉")

	animal = new(Dog)
	animal.eat("肉肉")
	// 人吃肉肉
	// 狗舔肉肉
}

// 接口就是一种规范,实现了这个接口的类必须实现接口的所有函数(这个类有接口的所有特性)
// 对于特定对象采用特定函数,和C的抽象类差不多