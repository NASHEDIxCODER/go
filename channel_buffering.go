package main

import "fmt"

/* by default chanels are in buffered, meaning that they will only accept sends (chan <-)
if there is  a corresponding receive (<- chan) ready to recieve the sent value.
buffered channels accept a limited number of values without a corresponding reciever for those values.
*/


func main() {
	messages := make (chan string,2) // we make a channel of string buffering up to 2 values.

// because this channenl is buffered, we can send these values into the channels without a corresponding concurrent receive.

	messages <- "buffered"
	messages <- "channel"
	// messages <- "channel2"  // this will give  error eg. all go routines are asleep - deadlock


	fmt.Println(<-messages) // we can recive two values as usual
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
}