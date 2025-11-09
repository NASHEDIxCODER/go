package main

import (
	"flag"
	"fmt"
)

//cli arguments with flag for example lablk -f here -f is flag

//go provides a flag package supporting basic command-line flag parsing.


func main() {
	
	//basic flg declearation are available for string integer and boolean options. here we decleare a string flag word.
	// with a default value "foo" and short description. this flag.String function returns a string pointer.


	wordPtr := flag.String("word","foo","a string")

// this declares numb and fork flags.

	forkPtr := flag.Bool("fork",false,"a bool")
	numPtr :=flag.Int("numb",42,"an int")

// it's also possible to declare an option an option that use an existing var declared elsewhere in the program 

	var svar string
	flag.StringVar(&svar, "svar","bar","a string var")
// once all flags are declared call flag.Parse() to execute the comman-line parsing.
	flag.Parse()
	
	fmt.Println("word:",*wordPtr)
	fmt.Println("numb:",*numPtr)
	fmt.Println("fork:",*forkPtr)
	fmt.Println("svar:",svar)
	fmt.Println("tail:",flag.Args())
}