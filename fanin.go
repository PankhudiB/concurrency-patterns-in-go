package main

// fan-in
func fanin(ch1, ch2 <-chan string) <-chan string {
	newChannel := make(chan string)

	go func() {for {newChannel <- <-ch1}}()
	go func() {for {newChannel <- <-ch2}}()

	return newChannel
}

