package main


type Person struct{

}

type People Person
// type People = Person		报错: ambiguous selector A.Say 上面的就不会
func (p Person) Say() {

}
type Student struct {
	Person
	People
}
func main() {
	A := Student{}
	A.Say()
}
