package main

import (
	"fmt"
	"pattern"
)

func main() {
	ksf := pattern.KissmeFactory{}
	bkf := pattern.BakechaFactory{}

	fmt.Println(ksf.CreateFood(), ksf.CreateCar())
	fmt.Println(bkf.CreateFood(), bkf.CreateCar())

	// {} {}
	// {} {}
}
