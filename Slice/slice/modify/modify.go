package modify

import "fmt"

//ACCESS elements of a slice
var mySLICE []int = []int{1,2,3,4,5,6,7,8,9}

func access(){
	fmt.Println(mySLICE[0])
	fmt.Println(mySLICE[len(mySLICE)-1]) // the last element
}

func Modify(){
	access()
}