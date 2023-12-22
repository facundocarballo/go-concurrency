package selectstatement

func Select() {
	chInt := make(chan int)
	chFloat := make(chan float64)
	chStr := make(chan string)

	ExecuteProducers(chInt, chFloat, chStr)
	ExecuteConsumers(chInt, chFloat, chStr)
}
