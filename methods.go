package main
//go supports methords defined on struct types.

import "fmt"

type rect struct {
	width, height int
}
func (r *rect) area() int{ // this area method has a receiver type of *rect.
	return r.width * r.height
}

func (r rect) perim() int { //methods can be defined for either pointer receiver types
	return 2 *r.width + 2*r.height
}

func main() {
	r := rect{width: 20, height: 10}

	fmt.Println("area: ", r.area()) //here is two method call defined in struct
	fmt.Println("perim: ",r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("parim: ", rp.perim())

}//go automatically handles conversion between values and pointers for method call.
//use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
