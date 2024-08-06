package arrays

import "fmt"


var x [5]int 
func Array(){

	for index,_ := range x{
		
		x[index]= index+1-1
		fmt.Printf("index[%d] is %d \n",index,x[index+1-1])
	}

	

}

func FixedArray(){
	x:=[10]float64{99,98,96,7,5,2,5,5,}

	for index,value :=range x{
		fmt.Printf("index[%d] is %f \n",index,value)

	}
}

func MultiDimension(){
	x:= [4][2]int{{1,1},{2,4},{3,9},{4,16},}

	for  _,value :=range x{
		//return key value
		fmt.Printf("value %d",value)
		
	}

	for i:=0;i < len(x);i++{
		for j:=0;j<len(x[i]);j++{
			fmt.Printf("index[%d][%d] is %d \n",i,j,x[i][j])

	}}
}