package problems

import (
	"fmt"
	"sync"
	"time"
)

var value = 0
var mutexValue sync.Mutex

func GoroutineC() {
	mutexValue.Lock()
	fmt.Printf("[GoroutineC] value++")
	value++
	mutexValue.Unlock()
}

func Starvation() {
	mutexValue.Lock()
	go GoroutineC()

	for {
		fmt.Printf("Starvation...\n")
		time.Sleep(time.Second * 3)
	}
}
