package main

import (
	"fmt"
	"time"
)

type handlerFunctionType = func(parameters ...interface{})
type middlewareType = func(handlerFunctionType) handlerFunctionType
type Handler struct {
	handlerFunction        handlerFunctionType
	middlewares            []middlewareType
	finallyHandlerFunction handlerFunctionType
}

func (h *Handler) SetHandlerFunction(handlerFunction handlerFunctionType) {
	h.handlerFunction = handlerFunction
}
func (h *Handler) AddMiddleware(middleware middlewareType) {
	h.middlewares = append(h.middlewares, middleware)
}
func (h *Handler) format() {
	h.finallyHandlerFunction = h.handlerFunction
	for i := 0; i < len(h.middlewares); i++ {
		h.finallyHandlerFunction = h.middlewares[i](h.finallyHandlerFunction)
	}
}

func (h *Handler) Do(parameters ...interface{}) {
	// 为空时先格式化
	if h.finallyHandlerFunction == nil {
		h.format()
	}
	// 这里可能还为空，所以还要判断下
	if h.finallyHandlerFunction != nil{
		h.finallyHandlerFunction(parameters...)
	}
}


// 处理器
func handlerFunction(parameters ...interface{}) {
	for i := 0; i < 100000000; i++ {}
}

// 中间件 (输出函数执行)
func countSpendTime(handler handlerFunctionType) handlerFunctionType {
	return func(parameters ...interface{}) {
		before := time.Now()
		handler(parameters...)
		fmt.Println(time.Since(before))
	}
}

// 中间件 (修饰)
func decorate(handler handlerFunctionType) handlerFunctionType {
	return func(parameters ...interface{}) {
		fmt.Println("-----------       ------------")
		handler(parameters...)
		fmt.Println("-----------   *   ------------")
		fmt.Println()
	}
}






func main() {
	handler := Handler{}
	handler.SetHandlerFunction(handlerFunction)
	handler.AddMiddleware(countSpendTime)
	handler.AddMiddleware(decorate)
	handler.Do()
}
