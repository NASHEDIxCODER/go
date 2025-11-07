package main
// interfaces are named collections of method signatures.
//here's is basic interfaces for geometric shapes.

import(
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type rect struct { // implemet this interface on rect and circle types
	width, height float64
}
type circle struct { 
	radius float64
}

func (r rect) area() float64{ // to implement an interface in go, we just need to implement all methodsin the interface.
	return r.width *r.height //we implement geometry in rects.
}
func (r rect) perim() float64{
	return  2*r.height +2*r.width
}

func (c circle) area() float64{ // implementation for circles.
	return 2 * math.Pi * c.radius *c.radius
}
func (c circle) perim() float64{
	return 2* math.Pi *c.radius
}
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle (g geometry){ //sometimes it's useful to know the runtime type of of an interface value.
	if c, ok := g.(circle); ok { //one option is using a type  assertion as shown here another is a type swtich.
		fmt.Println("circle with radius:",c.radius)
	}
}
func main(){

	r:= rect{width: 3,height: 4}
	c:= circle{radius: 5}
	
	measure(r) //the circle and rect struct types both implement the geometry interfaces so we can use instances of these structs as arguments to measure.
	measure(c)
	
	detectCircle(r)
	detectCircle(c)
	

}


