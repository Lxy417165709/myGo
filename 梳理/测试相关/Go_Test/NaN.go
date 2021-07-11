package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.0
	b:=(1.0-1.0)
	fmt.Println(a/b)
	fmt.Println(math.IsNaN(a/b))
}
