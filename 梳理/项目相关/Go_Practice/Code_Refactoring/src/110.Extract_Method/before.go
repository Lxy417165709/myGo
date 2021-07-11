package main

import "fmt"

// 重构前
type Test struct {
	name     string
	elements []float64
}

func (t *Test) printOwing(previousAmount float64) {
	outstanding := previousAmount * 1.2

	// print banner
	fmt.Println("***************************")
	fmt.Println("****** Customer Owes ******")
	fmt.Println("***************************")

	// calculate outstanding
	for _, v := range t.elements {
		outstanding += v
	}

	// print details
	fmt.Printf("%s: %s\n", "name", t.name)
	fmt.Printf("%s: %f\n", "amount", outstanding)
}

func main() {
	t := Test{"hello",[]float64{1,2,3,4,5}}
	t.printOwing(100)
}
