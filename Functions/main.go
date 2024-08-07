package main

import (
	"fmt"
	"gofunctions/data"
	"gofunctions/function"
	
)


func main(){
	
	fmt.Println("The average",function.ComputeAverage(data.XX))
	average,total := function.ComputeAverageAndsum(data.XX)
	fmt.Println("Average and sum",average,total)
	
}