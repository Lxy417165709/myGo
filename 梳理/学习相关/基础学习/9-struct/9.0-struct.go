package main

import "fmt"

// 定义结构体
type Student struct {
	id int
	name string
	sex string
}

func main(){
	// 初始化对象方法1
	studentA := Student{1,"李学悦","男"}
	// 初始化对象方法2
	studentB := Student{name:"Lxy",id:10,sex:"女"}
	fmt.Println(studentA, studentB)

	fmt.Println(studentB.id,studentB.name,studentB.sex)
	studentB.id = 100
	fmt.Println(studentB.id,studentB.name,studentB.sex)

}

// 访问结构体成员以及修改成员数据和C一样