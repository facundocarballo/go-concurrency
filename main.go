package main

import (
	"github.com/facundocarballo/go-concurrency/synchronization/selectstatement"
)

func main() {
	// goroutines.WithConcurrency()
	// synchronization.WithoutChannels()
	selectstatement.Select()
}
