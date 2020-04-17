package pattern


type Pearl struct {
	Drink Drink
	Price int
}
func (p Pearl)CountPrice()int{
	return p.Price + p.Drink.CountPrice()
}
