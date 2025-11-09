package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	//to set a key value pair use os.Setenv. to get a value of a key use os.Getenv. this will return an empty stringif the key isn't present.

	os.Setenv("FOO","1")
	fmt.Println("fOO", os.Getenv("FOO"))
	fmt.Println("Bar:", os.Getenv("BAR"))

	// use os.getEnviron to list all key / value pair in environment.

	//this returns a slice of strings in the form Key=value. you can strings.SplitN them to get the any key and value.

	fmt.Println()
	for _, e:= range os.Environ(){
		pair:= strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}


	//check for USER if presents then prints if not then pain occured.
	user, err:= os.LookupEnv("USER")
	if !err{
		panic("Environment variable 'USER' not found!")
	}
	fmt.Println("USER =",user)

}
