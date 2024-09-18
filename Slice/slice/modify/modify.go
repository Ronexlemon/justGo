package modify

import "fmt"

//ACCESS elements of a slice
var mySLICE []int = []int{1,2,3,4,5,6,7,8,9}

func access(){
	fmt.Println(mySLICE[0])
	fmt.Println(mySLICE[len(mySLICE)-1]) // the last element
}

//CHANGE element of a slice
func change(){
	for i:=0;i < len(mySLICE);i++{
		mySLICE[i]= mySLICE[i]+1

	}
	
	fmt.Println(mySLICE)
}

//APPEND elements to a slice
//SYNTAX slice_name= append(slice_name,element1,element2,...)
func apppend(){
	
		mySLICE = append(mySLICE,10 )
		fmt.Println(mySLICE)
		//append slice to slice
		mySLICE = append(mySLICE, mySLICE...)
		fmt.Println(mySLICE)
	
}


func Modify(){
	access()
	change()
	apppend()
}