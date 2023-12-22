package selectstatement

import "fmt"

func PrintInt(n int) {
	fmt.Printf("[PrintInt] Read %d\n", n)
}

func PrintFloat(n float64) {
	fmt.Printf("[PrintFloat] Read %.2f\n", n)
}

func PrintString(n string) {
	fmt.Printf("[PrintString] Read %s\n", n)
}

func ExecuteConsumers(
	chInt chan int,
	chFloat chan float64,
	chStr chan string,
) {
	for {
		select {
		case data := <-chInt:
			PrintInt(data)
		case data := <-chFloat:
			PrintFloat(data)
		case data := <-chStr:
			PrintString(data)
		}
	}
}
