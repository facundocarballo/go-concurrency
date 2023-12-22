package synchronization

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Producer(ch chan int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(time.Second * 3)
	}
}

func Consumer(ch chan int) {
	for {
		n := <-ch
		fmt.Printf("[Consumer] Read %d\n", n)
	}
}

func WithChannels() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(2)

	go func(channel chan int) {
		defer wg.Done()
		Producer(channel)
	}(ch)

	go func(channel chan int) {
		defer wg.Done()
		Consumer(channel)
	}(ch)

	wg.Wait()
}

//

func P(resource *int) {
	for {
		*resource = rand.Intn(100)
		time.Sleep(time.Second * 3)
	}
}

func C(resource *int) {
	for {
		fmt.Printf("[Consumer] Read %d\n", *resource)
		time.Sleep(time.Second * 2)
	}
}

func WithoutChannels() {
	var wg sync.WaitGroup
	resource := 0
	wg.Add(2)

	go func(resource *int) {
		defer wg.Done()
		P(resource)
	}(&resource)

	go func(resource *int) {
		defer wg.Done()
		C(resource)
	}(&resource)

	wg.Wait()
}
