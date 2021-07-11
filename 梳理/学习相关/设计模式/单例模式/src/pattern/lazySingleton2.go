package pattern

import "sync"

// 懒汉式
// 采用 doOnce
type LazySingleton2 struct {
}

var lazyInstance2 *LazySingleton2
var once sync.Once

func GetLazySingleton2() *LazySingleton2 {

	once.Do(func() {
		lazyInstance2 = &LazySingleton2{}
	})
	return lazyInstance2
}
