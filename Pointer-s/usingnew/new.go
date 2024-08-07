package usingnew

import "fmt"

func swap(xPtr *int,yPtr *int){
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp

}

func DoSwapWithNew(){
	xPtr := new(int)
	yPtr := new(int)
	*xPtr =10
	*yPtr =30

	swap(xPtr,yPtr)
	fmt.Printf("x before swap was 10 and y was 30 after swap x is %d and y is %d\n",*xPtr,*yPtr)

}