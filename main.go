package main

import "fmt"

func main() {
	channel1 := generator("a")
	for i := 0; i < 5; i++ {
		fmt.Println("Received on channel : ", <-channel1)
	}

	channel2 := generatorDataArray([]string{"a", "b", "c", "d"})
	for i := 0; i < 5; i++ {
		fmt.Println("Received on channel : ", <-channel2)
	}

}
