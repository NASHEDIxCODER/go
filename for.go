package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}
	var n int = 1
	for {
		fmt.Println("loop", n)
		n = n + 1
		if n > 100000000 {
			break
		}
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

//basic for loop has three components seprated by semicolons: init statement, conditions expression, and post element statement.
// the init statement is executed before the first iteration.
//the condition expression is evaluated before each itreation. if the evaluates to false , the loop ends.
// the post element statements is executed at the end of each iteration.
//the init and post statements are optional.
//if the condition is omitted, it is equivalent to true, creating an infinite  loop.
//the loop can be exited with a break statement.
// the continue statement skips to the next iteration of the loop.
// a for loop can also be written as a while loop, with only a condition.
// Go's for loop can also iterate over elements of a collection such as an array or slice using range keyword.
// the range from can be used to iterate over keys and values of a map or characters of a string as well.
//when ranging over an array, slice, string, or map, two values are returned and can be assigned to variables.
//the first value is the index or key, and the second is a copy of the element at the index or key.
