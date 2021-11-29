## What is Concurreny?
- Concurreny is about multiple thinks happening at the same time.
- Go provides built in support for concurreny.

## Why we need to think about Concurrency?
```
func Add(numbers[]int) int64 {
    var sum int64
    for _, number := range numbers {
        sum += int64(number)
    }
    return sum
}
```
If we have millions of numbers, it's gonna take a lot of time to add them up. 
When we run to the function it's running in one core and the other cores are idle.
Our target is to run the function in multiple cores.

- Concurrency is composition of independent execution computations, which may or may not run in parallel.
- Parallelism is the ability to execute multiple computations in simultaneous.
- Concurrency enables parallelism.

## What is a Process?

- An instance of a running program.
- Process provides environment for running a program.

|Stacks|
|------|
|Heap|
|Data|
|Code|

- OS -> Allocate memory.
- Code -> Machine instructions.
- Data -> Global data
- Stacks -> Used to store local variables.
- Heap -> Dynamic memory allocation.

## What is a Thread?
- Threads are the smallest units of execution.
- Process has at least one thread main thread
- Threads share the same memory space.
- Processes can have multiple threads.

|Heap|
|------|
|Data|
|Code|
|Thread-1 ~ Thread-2| 
|Stack ~ Stack|
|Registers ~ Registers|
|PC~PC|

- Threads run independently of each other.
- Threads share the same memory space.
- Threads can run concurrently in parallel.

## Thread States
- When the process is created, the main thread is put into the ready queue. It's in the runnable state.
- Once the CPU is available, the thread starts to execute and each thread given a time slice.
- If that time slice is over, the thread is put back into the ready queue.
- If the thread is blocked, it's put into the blocked queue.

Runnable --Scheduler Dispatch--> Executing --I/O or event wait--> Waiting --I/O or event completion--> Runnable

## C10k Problem
The C10k problem is the problem of optimizing network sockets to handle a large number of clients at the same time.
![alt text](http://monkey.org/~provos/libevent/libevent-benchmark.jpg)

- OS gives a fixed stack size for each thread. It's dependent on the hardware. So if I have a 8GB of memory and 8k kbytes stack, then theoretically I create 1000 thread.
 you can check with this command  `ulimit -a`

Go's Concurrency Tool Set
- Goroutines -> Goroutines are concurrently executing functions.
- Channels -> Channels are used to communicate between goroutines.
- Select -> Select is used to multiplex the channels.
- Sync -> Sync is used to synchronize the execution of goroutines.

## Goroutines
 - Goroutines extremely lightweight.
 - Starts with 2kb of stack, which grows and shrinks as needed.
 - Low CPU overhead.
 - Channels are used for the communication of data between goroutines. Sharing of memory can be avoided by using channels.
 - Go runtime creates worker OS threads.
 - Goroutines run in the context of OS thread.

## M:N Scheduling
- Go Scheduler runs in user space.
- Go Scheduler runs in the context of OS thread.
- Go runtime create number of worker OS threads, equal to the number of CPUs (GOMAXPROCS).
- Go Scheduler distributes runnable goroutines over multiple OS threads.

## Channels
- Channels are used to communicate between goroutines.
- Sync with Goroutines.
- Typed.
- Thread safe.
- example-1: `var ch chan <Type> && ch = make(chan <Type>)`
- example-2:`ch := make(chan <Type>)` allocate memory for channel.
- Pointer operator is used for sending and receiving the value from channel. The arrow indicates the direction of the communication.
  - Send: `ch <- <value>`
  - Receive: `<value> = <- ch`
  - Close: `Close(ch)` close the channel.
- Receive returns two value, the first one is a received value from the channel and the second one is a boolean value.
  - If the channel is closed, the second value will be false.
  - If the channel is not closed, the second value will be true.

### Range Channel
 - Iterate over values received from a channel.
 - Loop automatically stops when the channel is closed.
 - Range does not return a second boolean value.

### Unbuffered Channel
 - `make(chan <Type>)` allocate memory for channel.
### Buffered Channel
 - `make(chan <Type>, <Size>)` allocate memory for channel with buffer size.
 - in-memory FIFO queue.
 - Asynchronous.
### Ownership of channel avoids
 - Deadlocking by writing to nil channel.
 - Closing a nil channel.
 - Writing to a closed channel.
 - Closing a channel more than one

These are the reasons getting `panic`.

## Select
```
select {
case <-ch1:
    // do something
case <-ch2:
    // do something
default:
    // do something
}
```
- Empty select block will block forever.
- Select is like this switch statement.
- Select will block until any of the case is ready.
- With select we can implement a non-blocking communication and timeout.
- Select on nil channel will block forever.