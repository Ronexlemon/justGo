package pointers

import "fmt"

func Zero(xPtr *int){
	*xPtr = 0  //dereference and assign
}

func AssignZero(){
	x:=5
	xPtr := &x
	Zero(xPtr)
	fmt.Println(x)
	fmt.Println(xPtr)

	
}

func swap(xPtr *int, yPtr *int){
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
	
}


func DoTheSwap(){
	x :=10
	y:=20

	xPtr:= &x
	yPtr := &y
	swap(xPtr, yPtr)
	fmt.Printf("x before was 10 and y was 20 now y is %d and x is %d ",*yPtr,*xPtr)

}