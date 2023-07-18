package main

import "fmt"

// takes two int channels, stores right val (+1) into left
func rightToLeft(left, right chan int) {
	left <- 1 + <-right // 1st right read, locks until left read
}

func daisy_chaining() {
	const n = 10000
	var channels [n + 1]chan int
	for i, _ := range channels {
		channels[i] = make(chan int)
	}
	for i := 0; i < n; i++ {
		go rightToLeft(channels[i], channels[i+1])
	}

	// insert a value into right-hand end
	go func(c chan<- int) { c <- 1 }(channels[n])

	// get value from the left-hand end
	fmt.Println(<-channels[0])
}
