package main

import "fmt"

func main() {
	// map声明
	var phone map[string]string
	fmt.Printf("%p\n",phone)
	// 0x0

	// map初始化
	phone = make(map[string]string)
	fmt.Printf("%p\n",phone)
	// 0xc04206a1e0

	// 增 (key不存在为增)
	phone["lxy"] = "123456789"
	phone["cgw"] = "123456789"
	fmt.Println(phone)
	// map[lxy:123456789 cgw:123456789]

	// 改 (key存在为改)
	phone["lxy"] = "987654321"
	fmt.Println(phone)
	// map[lxy:987654321 cgw:123456789]

	// 删 (被删除的key不存在不报错)
	delete(phone,"lxy")
	fmt.Println(phone)
	// map[cgw:123456789]

	// 查 (被查找的key不存在不报错,mark为真表示存在key,否则为不存在)
	value,mark := phone["cgw"]
	fmt.Println(value,mark)
	value,mark = phone["lxy"]
	fmt.Println(value,mark)
	// 123456789 true
	// false
}
