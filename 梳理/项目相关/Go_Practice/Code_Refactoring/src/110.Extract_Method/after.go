package main

import "fmt"

// 重构后
type Test struct {
	name     string
	elements []float64
}

func (t *Test) printOwing(previousAmount float64) {
	// print banner
	t.printBanner()

	// calculate outstanding
	outstanding := t.getOutstanding(previousAmount * 1.2)

	// print details
	t.printDetails(outstanding)
}

func (t *Test) getOutstanding(previousAmount float64) (result float64) {
	result = previousAmount
	for _, v := range t.elements {
		result += v
	}
	return
}
func (t *Test) printBanner() {
	fmt.Println("***************************")
	fmt.Println("****** Customer Owes ******")
	fmt.Println("***************************")
}
func (t *Test) printDetails(outstanding float64) {
	fmt.Printf("%s: %s\n", "name",t.name)
	fmt.Printf("%s: %f\n","amount", outstanding)
}

func main() {
	t := Test{"hello",[]float64{1,2,3,4,5}}
	t.printOwing(100)
}
