package main


// pointers allows to pass refrences to values and record within your program.
import "fmt"

// how pointers work in contrast to values with 2 functions: zeroval and zeroptr.zeroval has an int parameter,
//so arguments will be paseed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.

func zeroval(ival int ){
	ival =0
}

func zeroptr(iptr *int){
	*iptr = 0
}

func main(){
	i:= 1
	fmt.Println("initial:", i)


	zeroval(i) // function call
	fmt.Println("zeroval:",i)

	zeroptr((&i))//function call  the &i syntax gives the memory address of i, i.e. a pointer to i.
	
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)//pointer can be printed too.
	// zeroval doesn't change the i in main but zeroptr does 
	//beacause it has a refrence to the memory address for that variable 
	

}