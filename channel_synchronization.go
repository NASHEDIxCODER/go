package main

import (
	"fmt"
	"time"
)

/* we can use channels to synchronize execution across go routines.
here an example of using a blocking recieve to wait for a go routine to finish.
when waiting for multiple goroutines to finish, we perfer to use wait group.
*/

 func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //send value to notify done.
	// this is the function we'll run in a goroutine.
	// the done channel will be used to notify another goroutine that this function's work is done.

 }

func main() {
	// start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)

	go worker(done)
	<-done //block until we recive a notification from the worker on the channel.

	// if you removed the <- done line from this program,
	// the program could exit before the worker finished its work, or in some cases even before it started.
	
}
