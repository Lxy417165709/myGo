package main
type Fragment interface {
	a()
}
type GetPodAction struct {

}
func (g GetPodAction) a(){

}


func main() {
	var fragment Fragment = new(GetPodAction)
}