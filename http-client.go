package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
)

// the go standard library comes with excellent support for http client and servers in the net/http package.


func main() {
	//Requesting url via flag
	wordPtr :=flag.String("url","http://example.com"," URL to GET")
	flag.Parse()
	fmt.Println("Requesting:",*wordPtr)
	resp ,err := http.Get(*wordPtr)


	// resp, err := http.Get("https://google.com")

	if err !=nil{
		panic(err)
	}
	defer resp.Body.Close()

	//print http response body
	fmt.Println("response ststus:", resp.Status) 

	header:= resp.Header
	fmt.Println(header)
	scanner := bufio.NewScanner(resp.Body)

	//first five line of the response body
	for i := 0; scanner.Scan() && i < 5; i++{

		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err !=nil{
		panic(err)
	}

}