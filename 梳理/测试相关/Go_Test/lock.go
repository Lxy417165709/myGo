package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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

// 执行前需保证已 Format，否则将没有任何输出
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
func handlerOfMutex(parameters ...interface{}) {
	value,mutex,wg := int64(0),sync.Mutex{},sync.WaitGroup{}
	times,_ := parameters[0].(int)
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			value++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("加锁，最终value:", value)
}

// 处理器
func handlerOfAtomic(parameters ...interface{}) {
	value,wg := int64(0),sync.WaitGroup{}
	times,_ := parameters[0].(int)
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&value, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("原操，最终value:", value)
}

// 中间件 (输出函数执行)
func countSpendTime(handler handlerFunctionType) handlerFunctionType {
	return func(parameters ...interface{}) {
		before := time.Now()
		handler(parameters...)
		fmt.Println(time.Since(before))
		fmt.Println()
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	mHandler := Handler{}
	mHandler.AddMiddleware(countSpendTime)
	mHandler.SetHandlerFunction(handlerOfMutex)
	mHandler.Do(100000)
	mHandler.Do(1000000)

	aHandler := Handler{}
	aHandler.AddMiddleware(countSpendTime)
	aHandler.SetHandlerFunction(handlerOfAtomic)
	aHandler.Do(100000)
	aHandler.Do(1000000)

	// 设置：runtime.GOMAXPROCS(1)
	// 加锁，最终value: 100000
	// 156.5836ms
	//
	// 加锁，最终value: 1000000
	// 709.1394ms
	//
	// 原操，最终value: 100000
	// 94.7452ms
	//
	// 原操，最终value: 1000000
	// 571.0602ms

	// 设置：runtime.GOMAXPROCS(2)
	// 加锁，最终value: 100000
	// 27.9261ms
	//
	// 加锁，最终value: 1000000
	// 308.6829ms
	//
	// 原操，最终value: 100000
	// 31.9138ms
	//
	// 原操，最终value: 1000000
	// 337.0991ms

	// 总结：在单核CPU下，原操的效率高于加锁，但在多核CPU中，原操的操作一般低于加锁。
}
