package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup
	var shareList = make(map[string]string)
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		for len(shareList) == 0 {
			cond.Wait()
		}
		fmt.Println(shareList["rsc"])
		cond.L.Unlock()
		wg.Done()

	}()

	// this one writes changes to sharedRsc
	cond.L.Lock()
	shareList["rsc"] = "foo"
	cond.Broadcast()
	cond.L.Unlock()
	wg.Wait()
}
