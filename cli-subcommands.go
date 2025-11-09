package main

import (
	"flag"
	"fmt"
	"os"
)

//the flag package lets us easily define simple subcommands that have their own flag.

func main () {

	//we declare a subcommand using the NewFlagSet function and proceed to define new specific for this subcommand.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable",false,"enable")
	fooName := fooCmd.String("name","","name")

	//for a diffrent subcommand we can define diffrent supported flag.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level",0,"level")

	if len(os.Args)<2{
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)

	}
	switch os.Args[1] { //check subcommands is invoked
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("   enable:",*fooEnable)
		fmt.Println("   name:",*fooName)
		fmt.Println("   tail:", fooCmd.Args())
		
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("   level:",*barLevel)
		fmt.Println("   tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
		
	}

}