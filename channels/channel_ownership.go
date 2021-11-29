package main

import "fmt"

func main() {
	owner := func() <-chan int {
		ch:= make(chan int)
		go func() {
			defer close(ch)
			for i:=0; i<10; i++ {
				ch <- i
			}
		}()
		return ch
	}
	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Println("Received:", v)
		}
		fmt.Println("Done")
	}

	ch := owner()
	consumer(ch)
}
