package main

// fan-in
func fanin(ch1, ch2 <-chan string) <-chan string {
	newChannel := make(chan string)

	go func() {
		for {
			newChannel <- <-ch1
		}
	}()
	go func() {
		for {
			newChannel <- <-ch2
		}
	}()

	return newChannel
}

// fan-in using switch statement
// observe only 1 go routine instead of 2 like above

/*
	Select - is like a switch but with each case as communication
	ALL channels are evaluated
	selection blocks until one can proceed
	if multiple can proceed - then select chooses randomly
	a default clause executes immediately if no channel is ready
 */
func faninUsingSwitchStatement(ch1, ch2 <-chan string) <-chan string {
	newChannel := make(chan string)

	go func() {
		for {
			select {
			case s := <-ch1: newChannel <- s
			case s := <-ch2: newChannel <- s
			}
		}
	}()

	return newChannel
}
