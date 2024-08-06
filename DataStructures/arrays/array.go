package arrays

import "fmt"


var x [5]int 
func Array(){

	for val,index := range x{
		x[index]= val
		fmt.Printf("index[%d] is %d \n",index,val)
	}

	

}