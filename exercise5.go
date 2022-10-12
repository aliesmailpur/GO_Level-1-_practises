package main

import "fmt"

// at the package level scope, using the var keyword, create a variable
// with the indentifier y . the variable should be of the underlying type
// of your custom type x

// now use conversion to convert the type of the value stored in x
// to the underlying type



type ali int 
var x ali 

var y int


func main() {

    fmt.Println(x)
	fmt.Printf("%T\t" , x)
	x = 889
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\t", y)

}