package problems

import (
	"fmt"
	"sync"
	"time"
)

var resource = 0
var response = ""
var mu sync.Mutex

func DoWork(idx int) {
	mu.Lock()
	response += fmt.Sprintf("[Goroutine %d] Reads: %d\n", idx, resource)
	mu.Unlock()

	resource++

	mu.Lock()
	response += fmt.Sprintf("[Goroutine %d] Write: %d\n-----\n", idx, resource)
	mu.Unlock()
}

func RaceCondition() {
	for {
		initTime := time.Now()

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				DoWork(idx)
			}(i)
		}
		wg.Wait()
		if resource == 10 {
			response = ""
			resource = 0
			continue
		} else {
			finishTime := time.Now()
			println(response)
			fmt.Printf("[RaceCondition] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), resource)
			break
		}
	}
}
