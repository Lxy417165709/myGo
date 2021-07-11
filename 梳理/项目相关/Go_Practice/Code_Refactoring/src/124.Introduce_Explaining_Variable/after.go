package _24_Introduce_Explaining_Variable

import "strings"

type Object2 struct {
	platform string
	browser  string
	resize   int
}

func (o Object2) Foo() {
	isMac := strings.ToUpper(o.platform) == "MAC"
	isIE := strings.ToUpper(o.browser) == "IE"
	wasResized := o.resize > 0
	if isMac && isIE && wasResized {
		// doSomething
	}
}
