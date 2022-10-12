package main

import "fmt"

// create your own type. have the underlying type to be an int
// create a variable of your new type with the IDENTIFIRE "x" using the "var"
// print out the value of variable x
// print out the type of variable x
// assign 42 to variable x
// print out the value of x

type ali int 
var x ali 




func main() {

    fmt.Println(x)
	fmt.Printf("%T\t" , x)
	x = 42
	fmt.Println(x)

}