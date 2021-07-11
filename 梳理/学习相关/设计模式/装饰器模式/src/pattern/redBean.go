package pattern


type RedBean struct {
	Drink Drink
	Price int
}

func (r  RedBean)CountPrice()int{
	return r.Price + r.Drink.CountPrice()
}
