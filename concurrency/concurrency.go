package main

import (
	"fmt"
)

/*
First, we create a channel that will be common to all goroutines. Then we start 5 goroutines and pass the channel as an argument. Inside each goroutine, we write the gopher id to standard output as soon the gopher starts cooking the dish. We then send the gopher id from the goroutine back to the caller. From there, we’re back to the body of the main function where we receive the gopher id and record their finishing time.

Since we’re dealing with concurrent code, we lose the ability to predict the order of the output, however, we can observe how the channel blocks the execution, as a goroutine has to wait until the channel is available before it can send an id. One possible output is included below. Keep in mind that we’re probably using more goroutines than the number of cores on our machine, hence it’s likely that a single core is time-multiplexed to simulate the concurrency.
*/

func main() {
	c := make(chan int) // create a channel to pass ints

	totalCooks := 1000

	for i := 0; i < totalCooks; i++ {
		go cookingGopher(i, c) // start a goroutine
	}

	for i := 0; i < totalCooks; i++ {
		gopherID := <-c // Recieve a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // send a value back to main
}
