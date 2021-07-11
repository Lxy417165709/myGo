package _31_Remove_Assignments_to_Parameters

func discount(inputVal, quantity, yearToDate int) int {
	if inputVal > 50 {
		inputVal -= 2
	}
	if quantity > 100 {
		inputVal -= 1
	}
	if yearToDate > 10000 {
		inputVal -= 4
	}
	return inputVal
}
