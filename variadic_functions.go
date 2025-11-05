package main 

import "fmt"

func sum (nums ...int){
	fmt.Print(nums, " ")
	total := 0

	for _, num  := range nums {
		total += num 
	}
	fmt.Println(total)

}

func main(){
	sum(1,2)
	sum(1,2,3)
	num1 := []int{1,2,3,4}
	sum(num1...)

	fmt.Println("len:",len(num1))
}

//variadic functions can be called with any number of trailing arguments for example , fmt.Println is common variadic functions.
// here is a function that will take an arbitrary number of ints as arguments 

//within the functions the type of nums is equivalent to []int.
// we can call len nums , iterate over it with range, etc.
//variadic functionss can be called in the way with individual arguments.
