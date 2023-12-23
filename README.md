# Concurrency in Golang

## Goroutines
A Goroutine is a way to execute functions concurrently in Golang.
You can execute Goroutines in different ways:

```golang
    func DoWork() {
        // do something here
    }

    func main() {
        go DoWork()
    }
```

```golang
    func DoWork() {
        // do something here
    }

    func main() {
        go func() {
            DoWork()
        }()
    }
```

Executing functions concurrently attracts new problems that we have to figure out, above all Synchronization problems.

Because we don't know when it's executing the functions, we don't know in what order the functions will be executing.

## Synchronization

To synchronize the concurrent functions in Golang (Goroutines), we use Synchronization methods like
- Wait Groups
- Channels
- Select Statements
- Mutex
- Atomic

### Wait Groups
This Synchronization method allows you to wait for the Goroutines that you execute, so they can finish their work.
```golang
    func DoWork() {
        // do something here
    }

    func main() {
        var wg sync.WaitGroup

        wg.Add(1)
        go func() {
            defer wg.Done()
            DoWork()
        }()

        wg.Wait()
    }
```
#### wg.Add(1)
Here we are adding a new Goroutine to Wait
#### wg.Done()
Every time that a Goroutine executes this function, will reduce the number of Goroutines that we wait to finish.
This indicates that this Goroutine doesn't need to be waited anymore.
#### wg.Wait()
This function will be waiting for the number of Goroutines to get to zero.
Once the number of Goroutines gets to zero, will continue with the execution of the program.

### Channels
Channels allow Goroutines to communicate with each other, sending and receiving messages in an asynchronous way.
```golang
    var ch make(chan int)

    func Producer() {
        for {
            ch <- 1
            time.Sleep(time.Second * 3)
        }
    }

    func Consumer() {
        for {
            n := <- ch
            fmt.Printf("[Consumer] Read: %d\n", n)
        }
    }

    func main() {
        var wg sync.WaitGroup

        wg.Add(2)

        go func() {
            defer wg.Done()
            Producer()
        }()

        go func() {
            defer wg.Done()
            Consumer()
        }()

        wg.Wait()
    }
```
#### var ch make(chan int)
Here we create the channel, a channel of type integer. That means, that this channel only can store integers.

#### ch <- 1
Here the Producer function is storing the number 1 in the channel.

#### n := <- ch
Here the Consumer function is creating an integer variable, and storing there the first integer that the channel has.
> **Important** If the channel doesn't have any integer there, the Consumer function will be waiting until an integer cames.
We have to be careful with that because can cause deadlock.

### Select Statement
Select statements in Golang, allow you to handle the consumption of multiple channels easily.

```golang
    var chInt make(chan int)
    var chFloat make(chan float64)

    func ProducerInt() {
        for {
            ch <- 1
            time.Sleep(time.Second * 3)
        }
    }

    func ProducerFloat() {
        for {
            ch <- float64(1)
            time.Sleep(time.Second * 3)
        }
    }

    func Consumer() {
        for {
            select {
                case data := <- chInt
                    fmt.Printf("Integer received from chInt channel: %d\n", data)
                case data := <- chFloat
                    fmt.Printf("Float received from chInt channel: %.2f\n", data)
            }
        }
    }

    func main() {
        var wg sync.WaitGroup

        wg.Add(3)

        go func() {
            defer wg.Done()
            ProduceInt()
        }()

        go func() {
            defer wg.Done()
            ProduceFloat()
        }()

        go func() {
            defer wg.Done()
            Consumer()
        }()

        wg.Wait()
    }
```
Here we have two differents channels, an integer channel and a float channel.
Each channel has its own producer, but both share the same consumer.
The consumer is smart enough to identify from which channel the data cames and print the data in a correct format.

### Mutex
This synchronized mechanism allows you to indicate when a Goroutine has access to some shared resource.

```golang
    var mutex sync.Mutex

func DoWork(balance *int) {
	mutex.Lock()
	*balance++
	mutex.Unlock()
	time.Sleep(time.Second * 1)
}

func main() {
	initTime := time.Now()
	balance := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			MDoWork(&balance)
		}()
	}
	fmt.Printf("[main] waiting Goroutine...\n")
	wg.Wait()

	finishTime := time.Now()
	fmt.Printf("[main] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
}
```
Here the main function executes 10 Goroutines at the same time, and each Goroutine wants to access the resource **balance** and incremented by 1.
With the mutex, we make sure that only one Goroutine can have full access to this particular resource.
If we don't use the mutex, we could have race conditions between our Goroutines.

#### mutex.Lock()
Blocks the shared resource, to make sure that this Goroutine has exclusive access to this resource until this Goroutine unlocks the mutex.

#### mutex.Unlock()
Unlock the shared resource, to allow other Goroutines to access this resource safely.

### Atomic
It's very similar to the mutex mechanism.
Here we want to make sure that only one Goroutine can access a particular resource.

```golang
    func DoWork(balance *int64) {
	    atomic.AddInt64(balance, 1)
	    time.Sleep(time.Second * 1)
    }

    func main() {
        initTime := time.Now()
        balance := int64(0)

        var wg sync.WaitGroup
        for i := 0; i < 10; i++ {
            wg.Add(1)
            go func() {
                defer wg.Done()
                ADoWork(&balance)
            }()
        }
        fmt.Printf("[main] waiting Goroutines...\n")
        wg.Wait()

        finishTime := time.Now()
        fmt.Printf("[main] executed in %d seconds. Balance: %d\n", finishTime.Second()-initTime.Second(), balance)
    }
```
It's the same example as the Mutex, but using **Atomic**