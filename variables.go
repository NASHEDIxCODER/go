package main

import "fmt"

func main() {

	var a = " initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}

// var is used to declare variable even more than one
// multiple variables can be declare in one line.
// variable declared without a corresponding value is zero valued.
// the := syntax is a shorthand for declaring and initializing a variable e.g for var f string = "apple" in this case this syntax is only available inside functions.
