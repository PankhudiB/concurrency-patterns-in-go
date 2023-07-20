package main

import (
	"fmt"
	"sync"
)

func multiplexing(in1, in2 <-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	wg.Add(2)
	go processSingle(out, &wg, in1)
	go processSingle(out, &wg, in2)

	go func() {
		fmt.Println("wait start")
		wg.Wait()
		fmt.Println("wait ended")
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
