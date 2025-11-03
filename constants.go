package main

import (
	"fmt"
	"math"
)

const a string = "constant"

func main() {

	fmt.Println(a)
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

// `const` declares a coonst value.
//  A `const` statement can also appear inside a function body.
// Constant expressions performs arithmetic with arbitrary precession.
//  A numeric cant has no type until it's given one such as by an explicit contersation.
//  Here `d` is a "float64" because math.sin expects a "float64" argument.
// Constant can be character Strings or boolean value.
// Constants cannot be decleared using the := sysntax.
// Constants values are evaluated at compile time and thus are more efficient than variable for values that never change.
// Constant can be used to create enumerated constants using the `iota` identifier.
// `iota` represents successive integer constants 0, 1, 2 ... with each const line in a block.
