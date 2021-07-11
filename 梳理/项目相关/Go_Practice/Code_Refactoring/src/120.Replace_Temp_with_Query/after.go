package _20_Replace_Temp_with_Query

type Object2 struct {
	quantity  float64
	itemPrice float64
}

func (o Object2) getPrice() float64 {
	return o.getBasePrice() * o.getDiscountFactor()
}
func (o Object2) getBasePrice() float64 {
	return o.quantity * o.itemPrice
}
func (o Object2) getDiscountFactor() float64 {
	if o.getBasePrice() > 1000 {
		return 0.95
	}
	return 0.98
}
