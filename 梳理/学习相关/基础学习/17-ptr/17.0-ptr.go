package main

import "fmt"
type action interface {
	run()
	eat()
}

type Person struct {
	Name string
}
// p是否指针都可以~
func (p *Person) run(){
	fmt.Println(p.Name + "正在跑~")
}
func (p Person) eat(){
	fmt.Println(p.Name + "正在吃~")
}

type Dog struct{

}
func (p Dog) run(){
	fmt.Println( "狗正在跑~")
}
func (p Dog) eat(){
	fmt.Println( "狗正在吃~")
}

/*func (p *Person) run2(){
	fmt.Println(p.Name + "正在跑~")
}*/
func (p Person) setName(name string) {
	p.Name = name
}
func (p *Person) setName2(name string) {
	p.Name = name
}
func main() {
	p := Person{"lixueyue"}
	fmt.Println(p)
	p.setName("123456")
	fmt.Println(p)
	//{lixueyue}
	//{lixueyue}

	p.setName2("123456")
	fmt.Println(p)
	// {123456}

	p.run()
	p.eat()
	//123456正在跑~
	//123456正在吃~

	d :=Dog{}
	d.run()
	d.eat()
	//狗正在跑~
	//狗正在吃~

	// func (p *Person) run()
	// func (p *Person) eat()
	// 如果重写的函数是这样的话, p是指针,此时下面的应该写为: var a action = &Person{"lxy"}
	// func (p Person) run()
	// func (p Person) eat()
	// 如果重写的函数是这样的话, p是指针,此时下面的应该写为: var a action = Person{"lxy"}
	var a action = &Person{"lxy"}
	a.eat()
	a.run()
	//lxy正在吃~
	//lxy正在跑~


	// func (d *Dog) run()
	// func (d *Dog) eat()
	// 如果重写的函数是这样的话, p是指针,此时下面的应该写为: a = &Dog{}
	// func (d Dog) run()
	// func (d Dog) eat()
	// 如果重写的函数是这样的话, p是指针,此时下面的应该写为: a = Dog{}
	a = Dog{}
	a.eat()
	a.run()
	//狗正在吃~
	//狗正在跑~




}


/*
	总结
	1.  接口重写的时候,func (name *type) ... 和 func (name type) ...都可以
		其中*的为指针传递,另外一个是指传递(会创建一个新的对象),指针传递可以正在修改传入对象的值,值传递不可以
	2.  var <变量名> <接口名>
		假如是 var a action,当Dog类实现了action接口后,要把Dog类对象赋值给a,需不需要加&取址符号取决于Dog重写函数
		func (name *type)  如果都是*的,那就要加&,否则不用(如果重写函数有*和没有*的,那会报错,不能赋值)
	3.  如果一个指针a指向一个对象,那在go中要调用该对象的方法,写法是  <指针名>.<方法名>()	     a.run()
		而在C++中,写法是  <指针名>-><方法名>()	a->run()

*/