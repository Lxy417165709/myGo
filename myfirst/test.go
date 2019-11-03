// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"reflect"
// )
// // 小写私有 大写才会显示
// type Student struct {
// 	Name string
// 	Id int
// 	Mark [3]int
// }

// func main(){
// 	var a = &Student{
// 		Name:"Lixueyue",
// 		Id : 2,
// 		Mark : [3]int{100,100,100},
// 	}
// 	fmt.Println(a)
// 	fmt.Println(reflect.TypeOf(a))
// 	b,err := json.Marshal(a)
// 	if err != nil {
// 	 	fmt.Println("encoding faild")
// 	} else {
// 		fmt.Println("encoded data : ")
// 		fmt.Println(b)
// 		fmt.Println(reflect.TypeOf(b))
// 	}



// }
