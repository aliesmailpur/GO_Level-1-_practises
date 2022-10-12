package main

import "fmt"

// using the code from previous exercise
// assign the following values to the three variables
// for x assign 42
// for y assign "james bond"
// for z assign true
// use Springf to print all the values to one single string
// assign the returned value of type stirng using the short
// decleration to a variable with the identifier S



var x  int =   42
var y   string = "James bond"
var z bool = true


func main() {

   s:= fmt.Sprintf( "%v\t%v\t%v"  ,  x , y , z)
   fmt.Println(s)

}
