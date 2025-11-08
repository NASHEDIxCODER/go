package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// a line filter is a common type of program that reeads input on stdin process it and then prints some derived result to stdout.
// grep and seed are common line filters.
//go example that writte a capitalized version of all input text.


func main(){
	scanner :=bufio.NewScanner(os.Stdin)

	for scanner.Scan(){//text returns the current token here the next line,from input.
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err!=nil{ //check for error during scan end of line is not expected and not reported as an error.
		fmt.Println(os.Stderr,"error:",err)
		os.Exit(1)
	}

}