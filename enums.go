package main

// enumerated types (enums) are a special case of sum types.
//an enum is a type that has fixed number of possible values, each with distinct name.
//go doesn't have an enum type as a distinct language feature , but enums are simple to implement using existing language idioms.

import "fmt"

type ServerState int  //enum type ServerState has an underlying int type.

const ( //possible values for ServerState are defined as constants.
	StateIdle ServerState = iota // the Special keyword `iots` generates seccesive constant value automatically in case 0,1,2and so on
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{ //by implementeing the fmt.Stringer interface, value of ServerState can be printed out or converted to strings.
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "Error",
	StateRetrying: "retrying",
}
 // this can get cumbersome if there are many possible values.
 //in such cases the Stinger tool can be used in confunction with go:genrate to automate the process.

 func (ss ServerState)String() string{
	return stateName[ss]
 }

 func main(){ // if we have a value of type int, we cannot pass it to transition the compiler will complian about type mismatch.
	ns := transition(StateIdle) // this provides some degree of compile-time type safety for enums.
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)

 }

 func transition (s ServerState) ServerState{ // translation emulates a state transition for a server: it takes the existing state and return a new state.
	switch s{
	case StateIdle:
		return StateConnected
	case StateConnected,StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}

 }