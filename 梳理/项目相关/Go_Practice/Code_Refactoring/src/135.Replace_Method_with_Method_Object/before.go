package _35_Replace_Method_with_Method_Object

// 杜撰的例子
type Object1 struct {
}

func (o Object1) delta() int{
	return 100
}

func(o Object1)gama(inputVal,quantity,yearToDate int) int{
	importantValue1 := (inputVal*quantity) + o.delta()
	importantValue2 := (inputVal*yearToDate) + 100
	if yearToDate - importantValue1 > 100 {
		importantValue2 -= 20
	}
	importantValue3 :=importantValue2*7
	// and so on
	return importantValue3 - 2 * importantValue1
}
