package main 

import (
	"fmt"
	"time"
	// multiple imports can be added here
)


func main() {
	now := time.Now()
	fmt.Println("currentt time : ", now)

	i:= 2 
	fmt.Println("writte ",i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		//without default case
	}
	//fallthrough example

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weeekend")
	default:
		fmt.Println("It's a weekday")
		// default case example 
	}

	//fall through example
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
		// no condition example
	}

	//fallthrough example
	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("nashedi_X_coder")

}

//switch statement example in go 
// demonstrate various forms of switch statements including
// simple switch, switch with multiple cases, switch without conditions.
//type switch and default case usage 
//switch without an expression is and alternate way ti express if-else logic, Here we can also show how the case expressions can be non-constants.

//A type switch copmares types insted of values. you can use this to discover the type of interface value. in this exaplme t variable will have the type corresponding to it's clause.
