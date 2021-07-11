package _24_Introduce_Explaining_Variable

import "strings"

type Object3 struct {
	platform string
	browser  string
	resize   int
}

func (o Object3) Foo() {
	if o.isMac() && o.isIE() && o.wasResized() {
		// doSomething
	}
}
func (o Object3) isMac() bool{
	return strings.ToUpper(o.platform) == "MAC"
}
func (o Object3) isIE() bool{
	return strings.ToUpper(o.browser) == "IE"
}
func (o Object3) wasResized() bool{
	return o.resize > 0
}


