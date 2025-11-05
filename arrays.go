package main 
import "fmt"

func main() {
	var a [5] int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array a :", a)
	 // accessing elements of array 

	var b[5]int 
	fmt.Println("array b before initialization:",b)
	//before initialization all elements will be zero
	fmt.Println("len:",len(b)) //length of array
	fmt.Println("3rd element of b:",b[2]) //accessing 3rd element
	
	c := [...]int {1,2,3,4,5,6}
	//:= is used for implicit type declearation 

	fmt.Println("array c:",c) //printing array c
	fmt.Println("5th element of c: ",c[4])
	fmt.Println("len of c:",len(c))
	//dynamic array size

	twoD := [2][3]int{
		{1,2,3},
		{4,5,6},

	}
	fmt.Println("twoD array:", twoD)
	
	var twoD1 [2][3]int
	for i := range 2{
		for j:= range 3{
			twoD1[i][j] = i +j
		}
	}
	fmt.Println("twoD array after initialization:", twoD)

}

// we can set a value at an index using the array[index] = value sysntax, and get a value with array[index]
//builtin len() function returns length of array 
// we have compiler inferred array length using ... in array declearation.


// Output:
// [ nashedixcoder ~/go ]# go run arrays.go
// Array a : [1 2 3 4 5]
// array b before initialization: [0 0 0 0 0]
// len: 5
// 3rd element of b: 0
// array c: [1 2 3 4 5 6]
// 5th element of c:  5
// len of c: 6
// twoD array: [[1 2 3] [4 5 6]]
// twoD array after initialization: [[1 2 3] [4 5 6]]
