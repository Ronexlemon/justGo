package main

import (
	"fmt"
	problem "gofunctions/Problem"
	"gofunctions/data"
	"gofunctions/function"
	"gofunctions/internalfunctions"
	"gofunctions/recursion"
	"gofunctions/variadic"
)


func main(){
	
	fmt.Println("The average",function.ComputeAverage(data.XX))
	average,total := function.ComputeAverageAndsum(data.XX)
	fmt.Println("Average and sum",average,total)
	variadic.Add(1,2,3,4,5)
	fmt.Println(recursion.Factorial(5))

	defer internalfunctions.Second()
	internalfunctions.First()
	fmt.Println(problem.Fib(10))
	
}