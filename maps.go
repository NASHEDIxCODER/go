package main

import (
	"fmt"
	"maps"
)

func main(){

	m :=  make(map[string]int)

	//set key value pair using name[key] = value
	m["k1"] =7
	m["k2"] = 13
	fmt.Println("map:", m )

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"] //key doesnot existy
	fmt.Println("v3:", v3)
	fmt.Println("len: ", len(m))

	delete(m,"k2")
	fmt.Println("map: ",m)

	clear(m)
	fmt.Println("map after clear: ",m)

	_, prs := m["k2"] //prs will be false since k2 has been deleted
	fmt.Println("prs:", prs)

	n := map[string] int {"foo": 1, "bar": 2}
	fmt.Println("map :", n )

	n2 := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	if maps.Equal(n,n2){ //maps.Equal is used to compare two maps
		fmt.Println("maps are equal")
	}else{
		fmt.Println("maps are not equal")
	}

}
//make funtions is used to create a map
// note that maps appear in the form map[k:v k:v] when printed with `fmt.Println`.
// declare and initialize map in same line
//the builtin delete removes key/value pair from a map
//to remove all key/value pair from a map, use clear builtin function,
//if key doesnot exist in map delete is a no-op
//if the key doesn't exist then zero value of thr value type is returned.
//printing a map via fmt.Println shows all key/value pairs.
//maps contains a number of useful utility functios for maps.
//the optional second return value when getting values from a map indicates 
	//	if the key was present in the map this can be used to disambigute between missing key and keys with zero values like 0 or "". Here we didn't need the value itself,
	//so we ignored it with the blank identifier _.

//maps are Go'sbuilt-in associative datatype sometimes called hashes or dicts in other languages.

// //output:
// [ nashedixcoder ~/go ]# go run maps.go 
// map: map[k1:7 k2:13]
// v1: 7
// v3: 0
// len:  2
// map:  map[k1:7]
// map after clear:  map[]
// prs: false
// map : map[bar:2 foo:1]
// maps are not equal