package main

import (
	"fmt"
	"sync"
)

func example(size int) {
	for i := 0; i < size; i++ {
		fmt.Println(i)
	}
}


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		example(100)
	}()
	wg.Wait()
}
