# Concurrency in Golang

## Goroutines
A Goroutine is a way to execute functions concurrently in Golang.
You can execute Goroutines in differents ways:

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
This Synchronization method allows you to wait for the Goroutines that you execute, so they can finish his work.
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