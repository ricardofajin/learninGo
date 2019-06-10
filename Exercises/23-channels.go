package main

import "fmt"

/*
	Channels are the pipes that connect concurrent goroutines.
	You can send values into channels from one goroutine and receive those values into another goroutine.
*/

func main() {

	//Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine and receive those values into another goroutine.
	messages := make(chan string)

	//Send a value into a channel using the channel <- syntax.
	go func() { messages <- "ping" }() //Here we send "ping" to the messages channel we made above, from a new goroutine.

	//The <-channel syntax receives a value from the channel.
	msg := <-messages //Here we’ll receive the "ping" message we sent above and print it out.
	fmt.Println(msg)  //When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

	/*
		By default sends and receives block until both the sender and receiver are ready.
		This property allowed us to wait at the end of our program for the "ping" message
		without having to use any other synchronization.
	*/
}
