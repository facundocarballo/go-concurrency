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

	go GoroutineC()

	for {
		mutexValue.Lock()
		fmt.Printf("Starvation...\n")
		time.Sleep(time.Second * 3)
		mutexValue.Unlock()
	}
}
