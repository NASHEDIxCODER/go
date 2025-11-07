package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

 func main() {
	var name string

	fmt.Println("enter name:")
	fmt.Scan(&name) //reads only one string eg. if input will be "nashedi coder " it will read only nashedi and ignore everything after white space.

	fmt.Println("hello!",name) // white space will be  automatically added.
	var name2 string 
	fmt.Println("enter full name")
	fmt.Scanln(&name2) // still it reads only one string.
	fmt.Println("hello", name2)
	

	var name4 string  // this method reads whole line including white spaces.
	reader:= bufio.NewReader(os.Stdin) //reads standard input os.Stdin

	fmt.Println("enter name ")
	input,err := reader.ReadString('\n') //err is optional instead we can use blank identifier _.

	if err != nil{ //this if block is optional only for error 
		fmt.Println("Error reading input:", err) 
		return
	}

	name4 = strings.TrimSpace(input) // removes \n and stray spaces
	fmt.Println("hello",name4)


 }