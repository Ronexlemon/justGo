package pointers

import "fmt"

func Zero(xPtr *int){
	*xPtr = 0
}

func AssignZero(){
	x:=5
	xPtr := &x
	Zero(xPtr)
	fmt.Println(x)
	fmt.Println(xPtr)

	
}