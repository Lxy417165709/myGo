package _19_Inline_Temp

type Object1 struct {
}

func (o *Object1) isGreater() bool {
	basePrice := o.getBasePrice()
	return basePrice > 1000
}

func (o *Object1) getBasePrice() int {
	return 1000
}
