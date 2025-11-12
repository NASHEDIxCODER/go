package main

// A goroutine is a lightweight thread managed by the the run time.
// it allopws you to run functions con currently - that is, multiple tasks executing at the same time.




import (
	"fmt"
	"time"
)

func say (s string){
	for i := 0; i < 10 ; i++{
		fmt.Println(s)
		time.Sleep(500*time.Millisecond)
	}
}


func main() {
	go say("go routine")
	say("main") // try commenting thiis line ,we can see this program exits immediatitely because of main
	
	fmt.Println("this is goroutine")
}

//output may varies each run.
//both main and new goroutine run at the same time and the go runtime schedules them.

// when main() ends, all goroutines end too --the programs exits immediatly

