package main
//use range of some of the data structures we'have already learned.

import "fmt"

func main(){
	nums := []int {2,3,4}	//here we use range to sum the numbers in a slice. Arrays worklike this too.
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums{ //range on arrays and slices provides both the indes and value for each entry.
		if num == 3{    // above we didn't need the index 
			fmt.Println("index:",i) //so we ignored it with blank identifier _. 
		}	//sometimes we actually want the indexes though.
	}
	kvs := map[string]string{"a": "apple","b":"banana"} //range on the map iterates over key/value pairs.
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs{	//range can also iterate over just the keys of a map.
		fmt.Println("key :", k)
	}
	for i, c := range "go"{ //range on strings iterates over unicode points.
		fmt.Println(i, c) 		//the first value is the starting byte index of of the rune and the second rune itself.
	}
}


// range iterate over elements in a variety of built-in data structures. 