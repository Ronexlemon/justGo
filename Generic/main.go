package main

import "fmt"

func Add[T int | float64 | string](a T, b T)T{
	return  a+b
}


func main(){
	result := Add("ron","jon")
	fmt.Println(result)
}