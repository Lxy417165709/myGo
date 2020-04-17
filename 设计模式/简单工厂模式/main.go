package main

import (
	"fmt"
	"pattern"
)

func main() {
	factory := pattern.KissmeFactory{}

	naicha := factory.Create("naicha")
	snack := factory.Create("snack")
	other := factory.Create("other")
	fmt.Println(naicha, snack, other)

	// {} {} <nil>
}
