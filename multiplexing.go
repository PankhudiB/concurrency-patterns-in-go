package main

import (
	"fmt"
	"sync"
)

func multiplexing(in ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	wg.Add(len(in))

	for _, c := range in {
		go processSingle(out, &wg, c)
	}

	go func() {
		wg.Wait()
	}()
	return out
}

func processSingle(out chan string, wg *sync.WaitGroup, ch <-chan string) {
	for c := range ch {
		out <- c
	}
	fmt.Println("done with ch")
	wg.Done()
}
