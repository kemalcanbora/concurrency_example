package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}
	withdraw := func(amount int) {
		mu.Lock()
		balance -= amount
		mu.Unlock()
	}

	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			deposit(1)
			wg.Done()
		}
	}()

	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			withdraw(1)
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println(balance)
}
