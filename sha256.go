package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

//SHA256 hashes are frequently used to compute short identities for binary or text blobs.
//eg, TLS/SSL certificates use SHA256 to compute a certificate's signatures.
//compute SHA256 in GO.


func main(){
	var s string //variable to store string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter string for computing SHA256")

	input, err:= reader.ReadString('\n')
	if err != nil{
		fmt.Println("error reading input:",err)
		return
	}
	s =strings.TrimSpace(input) //store userinput into s varible
	
	

	h := sha256.New() //here we start hash
	h.Write([]byte(s)) //writte expects bytes. if you have string s, use []byte(s) to coerce it to bytes.
	bs := h.Sum(nil) // this gets the finalized hash results as a byte slice. the argument to sum can be used to append to and exixting byte slice: it usually not needed.
	fmt.Println(s)
	fmt.Printf("%x\n",bs)
	
}//other hashes also can be computed using same pattern.