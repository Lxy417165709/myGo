package main

import "pattern"

func main(){
	proxy := pattern.NewProxy(pattern.Subject{})
	proxy.DoSomething()

	// 我是代理~ 做之前我要做啥呢~
	// 做了一点小事情~
	// 我是代理~ 做之后呢~
}

// 代理就是在不修改被代理类的代码前提下，通过调用代理为其添加一些功能
