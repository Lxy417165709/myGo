package _20_Replace_Temp_with_Query

type Object1 struct {
	quantity  float64
	itemPrice float64
}

func (o Object1) getPrice() float64 {
	basePrice := o.quantity * o.itemPrice
	discountFactor := 0.0
	if basePrice > 1000 {
		discountFactor = 0.95
	} else {
		discountFactor = 0.98
	}
	return basePrice * discountFactor
}
