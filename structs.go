package main

//go structs are typed collection of fields.they are useful for grouping data togather to form records.

import "fmt"

// this person struct type has name and age fields.

type person struct {
	name string
	age  int
}

func newPerson(name string) *person { //newPerson constructs a new person structs with the given name.

	p := person{name: name} // go is a garbage collected language ; you can safely return a pointer to a local variable,
	//it will only be cleaned up by the garbage collector when there are no active refrences to it.
	p.age = 42
	return &p
}
func main() {
	fmt.Println(person{"bob", 20}) //this syntax creates a new struct.

	fmt.Println(person{name: "alice", age: 45}) // we can name the fields when initializing a struct.
	fmt.Println(person{name: "fred"})           //ommit field will be zero valued.
	fmt.Println(&person{name: "ann", age: 40})  // an &prefix yeilds a pointer to construct.
	fmt.Println(newPerson("jon"))               // it's idiomatic to encapsulate new struct creation in constructor.

	s := person{name: "sean", age: 50} //Access struct fields with a dot
	fmt.Println(s.name)

	sp := &s //we can also use dots with struct pointers - the pointers are automatically dereferenced.
	fmt.Println(s.age)

	sp.age = 51 //mutable structs
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println(dog)
}

//if a struct type is only used for a single value, we don't have ti give it a name
//the value have an anonymous struct type.
