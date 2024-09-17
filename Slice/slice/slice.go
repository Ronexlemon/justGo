package slice

import "fmt"

// create a slice with []datatype{values}

/**
*Syntax => slice_name := []datatype{values}
 */

var mySlice  []int =[]int{}

//The code above declaires an empty slice of 0 length and 0 capacity
//to initialze the slice during duration use
var mySlice2  []int =[]int{1,2,3,4,5}

//There are two functions that can be used to return the length and capacity of a slice
/**
*1. len()
*2. cap()
*/

//CREATE A SLICE FROM ARRAY
//Array syntax var myArray =[length]datatype{values}
//myslice = myArray[start:end]
var myArray = [5]int{1,2,3,4,5}
var mySlice3 []int= myArray[0:4]

func Slice(){
	fmt.Println(cap(mySlice2))
}