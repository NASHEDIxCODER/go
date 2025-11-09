package main

import (
	"fmt"
	"os"
)

//command line arguments are common wa to parameterize execution of programs.

func main(){

	//os.Args provides to raw command-line arguments.note that the first value in this slice is the path to the program.
	//and the os.Args[1:] holds th argument to the program.

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	//to experiment with cli arguments it's best to build a binary with go build first.
	
}