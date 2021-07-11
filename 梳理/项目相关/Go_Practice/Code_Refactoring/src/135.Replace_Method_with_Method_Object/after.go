package _35_Replace_Method_with_Method_Object

type Object2 struct {
}

func (o Object2) delta() int {
	return 100
}
func (o Object2) gama(inputVal, quantity, yearToDate int) int {
	// 委托
	g := NewGama(o, inputVal, quantity, yearToDate)
	return g.compute()
}

type Gama struct {
	object          Object2
	inputVal        int
	quantity        int
	yearToDate      int
	importantValue1 int
	importantValue2 int
	importantValue3 int
}

func (g Gama) compute() int {
	g.importantValue1 = (g.inputVal * g.quantity) + g.object.delta()
	g.importantValue2 = (g.inputVal * g.yearToDate) + 100

	// 分离
	g.importThing()

	g.importantValue3 = g.importantValue2 * 7
	return g.importantValue3 - 2*g.importantValue1

}

func (g Gama) importThing() {
	if g.yearToDate-g.importantValue1 > 100 {
		g.importantValue2 -= 20
	}
}


func NewGama(obj Object2, inputVal, quantity, yearToDate int) Gama {
	return Gama{object: obj, inputVal: inputVal, quantity: quantity, yearToDate: yearToDate}
}
