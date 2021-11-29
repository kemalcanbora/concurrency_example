package main

import (
	"fmt"
	"sync"
)

func wrongExample() {
	var wg sync.WaitGroup
	for i:=1; i<=3; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
// When the program is run, it prints:
// 4
// 4
// 4

// This is because by the time goroutine got the chance to run the value of "i" had already been incremented value 4.

func correctExample() {
	var wg sync.WaitGroup
	for i:=1; i<=3; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}


