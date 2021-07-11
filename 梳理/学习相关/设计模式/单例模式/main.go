package main

import (
	"fmt"
	"pattern"
)

func main() {

	a1 := pattern.GetLazySingleton1()
	a2 := pattern.GetLazySingleton1()
	a3 := pattern.GetLazySingleton1()
	fmt.Printf("%p %p %p\n", a1, a2, a3)

	b1 := pattern.GetLazySingleton2()
	b2 := pattern.GetLazySingleton2()
	b3 := pattern.GetLazySingleton2()
	fmt.Printf("%p %p %p\n", b1, b2, b3)

	c1 := pattern.GetHungerInstance()
	c2 := pattern.GetHungerInstance()
	c3 := pattern.GetHungerInstance()
	fmt.Printf("%p %p %p\n", c1, c2, c3)

	// 0xc042008098 0xc042008098 0xc042008098
	// 0xc0420080d0 0xc0420080d0 0xc0420080d0
	// 0x551d10 0x551d10 0x551d10

	// 以上表明每一组都是同一对象
}
