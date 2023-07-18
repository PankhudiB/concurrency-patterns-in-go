package main

import "fmt"

// fan-in that closes the output channel as soon as any one of the input channel closes
func faninWithGracefulClose(ch1, ch2 <-chan string) <-chan string {
	newChannel := make(chan string)

	go func() {
		for {
			data, ok := <-ch1
			if !ok {
				fmt.Println("input channel closed ..breaking ...")
				close(newChannel)
				break
			}
			newChannel <- data
		}
	}()
	go func() {
		for {
			data, ok := <-ch2
			if !ok {
				fmt.Println("input channel closed ..breaking ...")
				close(newChannel)
				break
			}
			newChannel <- data
		}
	}()

	return newChannel
}
