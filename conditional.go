package main 

import "fmt"

func main() {

	if 7%2==0{
		fmt.Println("7 is even ")
	}else{
		fmt.Println("7 is odd")
	}

	if 8%4==0{
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 ==0{
		fmt.Println("either 8 or 7 is even ")
	}

	if num :=9; num <0{
		fmt.Println(num, "is negative")
	}else if num < 10{
		fmt.Println("has one digit")
	}else{
		fmt.Println("has multiple digits")
	}
}

// Go's if statements allows you to execute code based on conditons.
// the conditions are boolean expressions that evaluate to true or false.
//if the condition is true , the code block follwing the if conditions is executed.
//if the condition is false , the code block os skipped, and any else if or else block is executed.
// you can chain multiple else if statements  to check for multiple conditions.
// you can also include an optional statement to execute before the conditions is evaluated.
//we can have if conditions without an else statement as well.
//the scope of any variable declared in the if statement is limited to the  if , elseif , and else blocks.
//logical operators such as && (and), || (or), and ! (not) can be usefull in if conditions to combine multiple boolean expressions.
