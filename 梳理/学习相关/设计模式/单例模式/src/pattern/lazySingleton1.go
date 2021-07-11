package pattern

import "sync"

// 懒汉式
// 采用 加锁 + 双重校验
type LazySingleton1 struct {
}

var lazyInstance1 *LazySingleton1
var mutexLock sync.Mutex

func GetLazySingleton1() *LazySingleton1 {
	if lazyInstance1 == nil {
		mutexLock.Lock()
		defer mutexLock.Unlock()
		if lazyInstance1 == nil {
			lazyInstance1 = &LazySingleton1{}
		}
	}
	return lazyInstance1
}
