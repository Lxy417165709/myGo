package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	fmt.Println("Locking (G0)")
	mutex.Lock()
	fmt.Println("Locked (G0)")
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("Locked (G%d)\n", i)
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("UnLocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock(G0)")
	mutex.Unlock()
	fmt.Println("UnLocked(G0)")

	wg.Wait()

	// Locking (G0)
	// Locked (G0)
	// Locking (G3)
	// Locking (G1)
	// Locking (G2)
	// ready unlock(G0)
	// UnLocked(G0)
	// Locked (G3)
	// UnLocked (G3)
	// Locked (G1)
	// UnLocked (G1)
	// Locked (G2)
	// UnLocked (G2)

	// 下面这种输出也是可能的
	// Locking (G0)
	// Locked (G0)
	// Locking (G3)
	// Locking (G1)
	// Locking (G2)
	// ready unlock(G0)
	// UnLocked(G0)
	// Locked (G3)
	// UnLocked (G3)
	// Locked (G1)  <----
	// Locked (G2)
	// UnLocked (G1)
	// UnLocked (G2)

}
