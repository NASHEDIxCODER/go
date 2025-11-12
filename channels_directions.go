package main

import "fmt"

/*when using channels as function parameter,
you can specify if a channel is meant to only send and receive values.
this specify increases the type safety of the program.
*/

func ping(pings chan <- string, msg string){ //this ping function only accepts a channel for sending values.
	//it would be compile-time error to try to recieve on this channel.
	pings <-msg
}

func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg 
	// the pong function accepts one channel for recieves (pings) and a second for sends(pongs).

}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

