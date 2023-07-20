package main

import (
	"fmt"
	"time"
)

func main() {
	//channel1 := generator("a", nil)
	//for i := 0; i < 5; i++ {
	//	fmt.Println("Received on channel : ", <-channel1)
	//}
	//
	//output := fanin(generator("a", nil), generator("b", nil))
	//for i := 0; i < 10; i++ {
	//	fmt.Println("Received on channel : ", <-output)
	//}

	//quit := make(chan bool)
	//output2 := faninUsingSwitchStatement(generator("c", quit), generator("d", quit), quit)
	//for i := 0; i < 10; i++ {
	//		data, ok := <-output2
	//		if !ok {
	//			fmt.Println("Someone closed channel..Exiting")
	//			break
	//		}
	//		fmt.Println("Received on channel : ", data, ok)
	//}

	//output3 := faninWithGracefulClose(generator("e", nil), generator("f", nil))
	//for i := 0; i < 10; i++ {
	//	data, ok := <-output3
	//	if !ok {
	//		fmt.Println("Someone closed channel..Exiting")
	//		break
	//	}
	//	fmt.Println("Received on channel : ", data, ok)
	//}

	//daisy_chaining()

	out := multiplexing(generator("a", nil), generator("b", nil))
	go func() {
		for i := 0; i < 10; i++ {
			data, ok := <-out
			if !ok {
				fmt.Println("Someone closed channel..Exiting")
				break
			}
			fmt.Println("Received on channel : ", data, ok)
		}
	}()
	time.Sleep(10 * time.Second)
}
