package synchronization

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func MDoWork(balance *int, whitMutext bool) {
	if whitMutext {
		mutex.Lock()
		*balance++
		mutex.Unlock()
	} else {
		*balance++
	}
	time.Sleep(time.Second * 1)
}

func WithMutex() {
	initTime := time.Now()
	balance := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			MDoWork(&balance, true)
		}()
	}
	fmt.Printf("[WithMutex] waiting Goroutines...\n")
	wg.Wait()

	finishTime := time.Now()
	fmt.Printf("[WithMutex] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
}

func WithoutMutex() {
	for {
		initTime := time.Now()
		balance := 0

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				MDoWork(&balance, false)
			}()
		}
		wg.Wait()
		if balance == 10 {
			continue
		}
		finishTime := time.Now()
		fmt.Printf("[WithoutMutex] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
	}
}
