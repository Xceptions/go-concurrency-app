package main

import "fmt"

func main() {
	// create a buffered channel - limited capacity
	// unbuffered channel waits for the receiver to 
	// receive the data before the sender gets unblocked
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}