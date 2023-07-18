package main

import "fmt"

func main() {
	channel1 := generator("a")
	for i := 0; i < 5; i++ {
		fmt.Println("Received on channel : ", <-channel1)
	}

	output := fanin(generator("a"), generator("b"))
	for i := 0; i < 20; i++ {
		fmt.Println("Received on channel : ", <-output)
	}
}
