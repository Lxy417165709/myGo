package pattern

import "fmt"

type Subject struct {
}

func (s Subject)DoSomething(){
	fmt.Println("做了一点小事情~")
}
