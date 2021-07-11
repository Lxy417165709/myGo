package _24_Introduce_Explaining_Variable

import "strings"

type Object1 struct {
	platform string
	browser string
	resize int
}

func (o Object1)Foo(){
	if strings.ToUpper(o.platform)=="MAC"  && strings.ToUpper(o.browser) == "IE" && o.resize > 0{
		// doSomething
	}
}



