package _31_Remove_Assignments_to_Parameters

func discount2(inputVal, quantity, yearToDate int) (result int) {
	result = inputVal
	if inputVal > 50 {
		result -= 2
	}
	if quantity > 100 {
		result -= 1
	}
	if yearToDate > 10000 {
		result -= 4
	}
	return
}
