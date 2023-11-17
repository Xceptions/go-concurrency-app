package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data in myChannel"
	}()

	go func() {
		anotherChannel <- "cow data in anotherChannel"
	}()

	// if multiple messages are available, it will pick one
	// at random
	select {
	case msgFromMyChannel := <- myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <- anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}