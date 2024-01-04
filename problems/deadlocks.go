package problems

import (
	"fmt"
	"sync"
)

var resourceA = 0
var resourceB = 0

var mutexA sync.Mutex
var mutexB sync.Mutex

func GoroutineA() {
	mutexA.Lock()
	mutexB.Lock()

	resourceA++
	resourceB++

	mutexA.Unlock()
	mutexB.Unlock()
}

func GoroutineB() {
	mutexA.Lock()
	mutexB.Lock()

	resourceA++
	resourceB++

	mutexA.Unlock()
	mutexB.Unlock()
}

func Deadlock() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		GoroutineA()
	}()

	go func() {
		defer wg.Done()
		GoroutineB()
	}()

	wg.Wait()
	fmt.Printf("ResourceA: %d\n", resourceA)
	fmt.Printf("ResourceB: %d\n", resourceB)
}
