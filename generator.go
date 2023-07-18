package main

import (
	"fmt"
	"time"
)

// generator returns a receive-only channel
func generator(data string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("%s-%d", data, i)
			time.Sleep(1 * time.Second)
		}
	}()
	return channel
}

func generatorThatClosesAfterNSeconds(data string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("%s-%d", data, i)
			time.Sleep(1 * time.Second)
			if i == 4 {
				close(channel)
				break
			}
		}
	}()
	return channel
}

func generatorDataArray(data []string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			for _, d := range data {
				channel <- d
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return channel
}
