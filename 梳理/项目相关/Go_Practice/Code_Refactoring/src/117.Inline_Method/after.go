package Inline_Method


type Object2 struct {
	number int
}

func (o *Object2) getRating() int {
	if o.number > 5 {
		return 2
	}
	return 1
}



