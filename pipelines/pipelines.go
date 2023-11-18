package main

import "fmt"

// create a channel, receive the input list
// into it. since I created an unbuffered channel,
// only 1 item can be placed into it until the value
// is read. This continues until the items in the list
// is used up
func sliceToChannel(in []int) <-chan int {
	out := make(chan int)

	go func(){
		for _, num := range in {
			out <- num
		}
		close(out)
	}()

	return out
}

// takes each item placed in the channel
// in sliceToChannel, square them, and
// place in the channel created in this
// function. Just as this treats the input
// from the sliceToChannel one entry at a
// time, the main method also treats its
// output one entry at a time.
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func(){
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main(){
	// input
	nums := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3

	for n := range finalChannel {
		fmt.Println(n)
	}
}