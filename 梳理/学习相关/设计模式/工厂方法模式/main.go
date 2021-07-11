package main

import (
	"fmt"
	"pattern"
)

func main() {
	ncf := pattern.NaiChaFactory{}
	sf := pattern.SnackFactory{}
	naicha := ncf.CreateFood()
	snack := sf.CreateFood()
	fmt.Println(naicha, snack)

	// {} {} <nil>
}
