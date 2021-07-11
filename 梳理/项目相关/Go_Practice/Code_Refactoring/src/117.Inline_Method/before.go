package Inline_Method


type Object1 struct {
	number int
}

func (o *Object1) getRating() int {
	if o.moreThanFiveLateDeliveries() {
		return 2
	}
	return 1
}
func (o *Object1) moreThanFiveLateDeliveries() bool {
	return o.number > 5
}


//func main(){
//	beginTest([]int{1,2,3,4,5,6,7,8,9,10})
//}
