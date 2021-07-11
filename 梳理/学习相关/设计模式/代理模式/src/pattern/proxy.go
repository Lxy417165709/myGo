package pattern

import "fmt"

type proxy struct {
	Subject Subject
}

func NewProxy(s Subject) proxy{
	return proxy{s}
}

func (p proxy)DoSomething() {
	p.beforeDo()
	p.Subject.DoSomething()
	p.afterDo()
}

func (p proxy) beforeDo(){
	fmt.Println("我是代理~ 做之前我要做啥呢~")
}
func (p proxy) afterDo(){
	fmt.Println("我是代理~ 做之后呢~")
}
