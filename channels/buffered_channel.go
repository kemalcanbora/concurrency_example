package main

import "fmt"

func main() {
	ch :=make(chan int, 6)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending %d\n ", i)
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Printf("Received %d\n", v)
	}
}
