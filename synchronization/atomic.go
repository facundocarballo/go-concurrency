package synchronization

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func ADoWork(balance *int64, whitMutext bool) {
	if whitMutext {
		atomic.AddInt64(balance, 1)
	} else {
		*balance++
	}
	time.Sleep(time.Second * 1)
}

func WithAtomic() {
	initTime := time.Now()
	balance := int64(0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ADoWork(&balance, true)
		}()
	}
	fmt.Printf("[WithAtomic] waiting Goroutines...\n")
	wg.Wait()

	finishTime := time.Now()
	fmt.Printf("[WithAtomic] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
}

func WithoutAtomic() {
	initTime := time.Now()
	balance := int64(0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ADoWork(&balance, false)
		}()
	}
	fmt.Printf("[WithoutAtomic] waiting Goroutines...\n")
	wg.Wait()

	finishTime := time.Now()
	fmt.Printf("[WithoutAtomic] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
}
