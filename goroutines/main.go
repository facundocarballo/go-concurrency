package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func DoWork(balance *int) {
	*balance++
	time.Sleep(time.Second * 1)
}

func WithoutConcurrency() {
	initTime := time.Now()
	balance := 0

	for i := 0; i < 10; i++ {
		DoWork(&balance)
	}

	finishTime := time.Now()
	fmt.Printf("[WithoutConcurrency] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
}

func WithConcurrency() {
	initTime := time.Now()
	balance := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DoWork(&balance)
		}()
	}
	fmt.Printf("[WithConcurrency] waiting Goroutines...\n")
	wg.Wait()

	finishTime := time.Now()
	fmt.Printf("[WithConcurrency] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
}
