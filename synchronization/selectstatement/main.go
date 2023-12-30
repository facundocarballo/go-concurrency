package selectstatement

import "sync"

func Select() {
	chInt := make(chan int)
	chFloat := make(chan float64)
	chStr := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		ExecuteProducers(chInt, chFloat, chStr)
	}()

	go func() {
		defer wg.Done()
		ExecuteConsumers(chInt, chFloat, chStr)
	}()

	wg.Wait()
}
