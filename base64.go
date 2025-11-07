package main

//go provides builtin support for base64 encoding and decoding.

import (
	b64 "encoding/base64"
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){
	var s string
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("enter a string to encode in base64") 

	input, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("error occured:",err)
		return
	}
	s = strings.TrimSpace(input)
	// fmt.Println(s)
	// above code handle user input only
	sEnc := b64.StdEncoding.EncodeToString(([]byte(s)))
	fmt.Println(sEnc)
// above code encode user input into base64

	sDec, _ :=b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	// sbove code decode the base64 string

	uEnc :=b64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(uEnc)
	//above code encode userinput into url encoding
	uDec ,_ :=b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))





	
}