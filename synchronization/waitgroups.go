package synchronization

import (
	"fmt"
	"sync"
	"time"
)

func DoWork(idx int) {
	time.Sleep(time.Second * 3)
	fmt.Printf("[Goroutine %d]\n", idx)
}

func WithWaitGroups() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			DoWork(idx)
		}(i)
	}
	wg.Wait()
	fmt.Printf("[WaitGroups] all goroutines finished.")
}

func WithoutWaitGroups() {
	for i := 0; i < 10; i++ {
		go DoWork(i)
	}
	fmt.Printf("[WithoutWaitGroups] finish.")
}
