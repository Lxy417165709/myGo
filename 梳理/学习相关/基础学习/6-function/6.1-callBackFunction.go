package main

import "fmt"

type cb func(int)

func callBack1(a int){
	fmt.Printf("我是回调函数1,我输出的值是%d\n",a)
}
func callBack2(a int){
	fmt.Printf("我是回调函数2,我输出的值是%d\n",a)
}
// 所谓回调,就是调用传入参数(该参数是函数类型)
func testCallBack(a int,fct cb){
	fct(a)
}
func main(){
	testCallBack(1, callBack1)
	testCallBack(2, callBack2)
}
