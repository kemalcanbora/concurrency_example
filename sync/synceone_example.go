package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	load := func() {
		fmt.Println("Run only one time!")
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			once.Do(load)
		}()
	}
	wg.Wait()
}
