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
	mutexB.Lock()
	mutexA.Lock()

	resourceA++
	resourceB++

	mutexB.Unlock()
	mutexA.Unlock()
}

func Deadlock() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			GoroutineA()
		}
	}()

	go func() {
		defer wg.Done()
		for {
			GoroutineB()
		}
	}()

	wg.Wait()
	fmt.Printf("ResourceA: %d\n", resourceA)
	fmt.Printf("ResourceB: %d\n", resourceB)
}
