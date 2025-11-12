package main

//channels are the pipe that connet concurrent go routines. 
//you can send values into channels from one goroutine and recive those values into another goroutine.
// it ensure synchronization between to channels.
import "fmt"

func main(){

	messages := make(chan string)

	go func() { messages <- "ping "} ()
// the <- channel syntax recives a value from the channel.
//here we'll recive the "ping" message we sent above and print it out.

	msg := <- messages 
	fmt.Println(msg)
	// when we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

}