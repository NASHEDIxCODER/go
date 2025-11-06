package main

import "fmt"

func intSeq() func () int {
	i:=0
	return func() int{
		i++
		return i
	}
}

func main(){

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}

//go supports anonymous functions which can form closures anonymous function are useful when you want to define a function inline without having to name it.
// this function intSeq return another function, which we define anonymous in the body of intSeq.
//the return function closes over the variable i to form closure.

//we call intSeq and assigning the result od a function to nexInt.
//this function value captures its own i value , which will be updated each time we call nextInt.
//to confirm that the state is unique to that particular function create and test a new one.
