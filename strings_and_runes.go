package main

// a go string is a read only slice of byte.
// the language and the standard library treat strings specially -as container of the text encoded in UTF-8.
// in other languages , string are made of "characters".
// in go, the concept character is called a rune - .
// it's an integer that represents a unicode  code point.

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "    a" // s is string assigned a literal value represents the word "hello " in the thai language.
	//go string literals are UTF-8 encoded text.
	//since striings are equivalent to zero  byte , this will producr the length of athe raw bytes stored within.
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ { //indexing into a string produces the raw byte value at each index.
		fmt.Printf("%x ", s[i]) //this loop genrates the hex value of all the bytes that constitute the code points in s.

	}
	fmt.Println()
	fmt.Println("rune count:", utf8.RuneCountInString(s)) // to count how many rune s in a strings , we can use the utf-8 package .
	//that the run time of RuneCountInString depends on the size of the string because it has to decode each UTF-* run sequentially.
	// some thai character are represented by utf-8 code points that can span multiple bytes.
	//so the results of this count may be surprising.

	for idx, runeValue := range s { //range loop handles strings specially and decodes each rune along with its offset in the string.
		fmt.Printf("%#U starts at %d\n", runeValue, idx)

	}
	fmt.Println("\n U?sing DecodeRuneInSrting ")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == ' ' {
		fmt.Println("found so sua")
	}
}
// values enclosed in single quotes are rune literals. we can compare a rune value to a rune directly

