package defers

import "fmt"

func  Defer(){
	welcome:= "Welcome to Defers go."
	//last in first out
	 defer fmt.Println(welcome)
	 defer fmt.Println("one")
	 defer fmt.Println("two")
	fmt.Println("Hello")
	MyDefer()
  	
}

func MyDefer(){
	for i:=0; i< 5;i++{
	defer	fmt.Print(i)
	}
}