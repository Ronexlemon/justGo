package pointers

import "fmt"

 func Pointer(){
	welcome:= "Welcome to the pointer lego"

	fmt.Println(welcome)
	//Declare a variable using var keyword
	// var ptr *int
	// fmt.Println("value of pointer is",ptr) //empty pointer is nill
	myNumber:=23

	var ptr = &myNumber
	fmt.Println("Value of actual pointer is:",ptr)
	fmt.Println("Value of actual pointer is:",*ptr)
	*ptr= *ptr+3
	fmt.Println("Value after addition is:",myNumber)
 }