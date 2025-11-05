package main 
import "fmt"

func vals() (int,int){
	return 4,8
}

func main (){
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)
	
	_, c := vals() // Ignoring the first return value
	fmt.Println(c)
	// fmt.Println(_) // This will cause a compile-time error
}

// the (int , int ) indicates that the function returns two integers.
// here we use 2 diffrent return values from the call with multiple assignment.
// if we want subset of the return values we can use blank identidier 