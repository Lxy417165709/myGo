package main

import (
	"encoding/json"
	"fmt"
)

// 如果结构体是这样的话
// 那json中的Sex字段存储的是struct中的Name,
// 那json中的Name字段存储的是struct中的Sex,
//type Person struct{
//	Name string 	`json:"Sex"`
//	Sex string		`json:"Name"`
//	Age int			`json:"Age"`
//}

// 这个json和struct中的数据才一一对应
type Person struct {
	Name string `json:"Name"`
	Sex  string `json:"Sex"`
	Age  int    `json:"Age"`
}

func main() {
	// 1. struct转json
	p := Person{"李学悦", "男", 10}
	b, _ := json.Marshal(p)
	js := string(b)
	fmt.Println(js)
	// {"Name":"李学悦","Sex":"男","Age":10}

	// 2. json转struct
	//p2 := Person{}	// 这个不可以...
	p2 := new(Person)
	json.Unmarshal([]byte(js), p2)
	fmt.Println(p2, string([]byte(js)), p2.Name)
	// &{李学悦 男 10} {"Name":"李学悦","Sex":"男","Age":10} 李学悦

}

/*
	总结
	1. struct转换json时要注意字段的首字母要大写,而且字段tag要写正确
	2. json转struct的时候,按的是struct字段的声明顺序,而不是tag
*/
