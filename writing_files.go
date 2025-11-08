package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//writing files in go is similar patterns.

func check(e error) {
	if e != nil {
		panic((e))
	}
}
func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("test/data.txt", d1, 0644) //0644 isthe file permission(an octal literal)used when creating file.
	check(err)

	f, err := os.Create("test/dat2.txt")

	defer f.Close() //its idiomatic a defer a close immediatly after opening a file.

	d2 := []byte{115, 111, 109, 101, 10} //creates byte slice manually using ascii character
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, err := f.WriteString("writte\n")
	check(err)
	fmt.Printf("wrote %d bytes \n", n3)
	f.Sync() //issue to flush writes in stable storage


	//bufio provides buffered writers in adition to a buffered reader we saw earlier.
	
	w := bufio.NewWriter((f)) //bufio provides buffered writers in addition to buffered readers we saw earlier.

	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() //use flush to ensure all bufferd operations have been applied to the underlying writer.

	//wwritte via user input.
	var s string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	check(err)
	s = strings.TrimSpace(input)
	// fmt.Println(s)
	z := bufio.NewWriter(f)
	n5, err := z.WriteString(s)
	check(err)
	fmt.Printf("wrote %d ", n5)
	z.Flush()

}
