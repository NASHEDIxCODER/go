package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("unit:", s, s==nil, len(s) ==0) // true true

	s = make([]string, 3) // len 3, cap 3
	fmt.Println("emp:",s , "len:", len(s), "cap:", cap(s))

	s[0] = "samrat"
	s[1] = "samrat2"
	s[3-1] = "nashedi" // 3-1 == 2 index
	fmt.Println("len:", len(s))
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	
	s = append(s, "samrat3")
	s = append(s, "coder")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))

	c := make([]string, len(s))// new slice c with len of s
	fmt.Println("cpy before: ",c , "len:", len(c))
	copy(c, s) // copy s to c
	fmt.Println("cpy:", c)
	l := s[2:5] //slicee from index 2 to 4
	fmt.Println("slice:", l)

	l = s[:5] // from start to index 4
	fmt.Println("slice2:", l)

	l = s[2:]
	fmt.Println("slice:",l)

	t := []string{"gopher", "samrat", "nashedi"}
	fmt.Println("dcl:", t)

	t2 := []string{"gopher","samrat", "nashedi"}
	fmt.Println("dclt2:",t2)
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	//twoD slices 
	twoD := make([][]int,3) //slices of slices
	for i := range 3{
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d slice:", twoD)


}


// we can declere and initialize slices in a single line.
// the slices package contains a number of useful utility funtion for slices.
// slices can be composed into multi-dimentional data structures.
//the length of the inner slices can vary, unlike with multi-dimentional arrays.
//slices can also be copy'd using the builtin copy function
// IsEmptySlice checks if a given slice is empty.
// Contains checks for nil slices and maps to avoid potential panics.