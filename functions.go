package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}
func greet(){
	fmt.Println("hello")
}

func main() {

	greet1 := func (name string)  { //anonymous function assigned to greet1 variable
		fmt.Println("hello ", name)
	}
	greet1("nashedi coder") //calling greet1 function
		
	greet() //calling greet function


	res := plus(10, 17)
	print("testing plus function ") //without fmt package can be use Print
	println("res:", res) //without fmt package can be use Println

	res2 := plusplus(2, 5, 7)
	fmt.Println("res2:", res2)//with fmt package use fmt.Println
}

//func is used to declare a function is go 
//function name is followed by parameter list enclosed in parenthesis.
// type comes after the parameter name.
//function cannot be nested within another function
//function can not be inside main function.
// however we can declare ananymous function inside main funtion or other functions 
