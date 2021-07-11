package _19_Inline_Temp

type Object2 struct {
}

func (o *Object2) isGreater() bool {
	return o.getBasePrice() > 1000
}

func (o *Object2) getBasePrice() int {
	return 1000
}
