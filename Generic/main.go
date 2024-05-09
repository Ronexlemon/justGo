package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// type num interface{
// 	int | float64 | string |float32 |int16 |int8
// }

// func Add[T num](a T, b T)T{
// 	return  a+b
// }

func Add[T constraints.Ordered](a T, b T)T{
	return  a+b
}

func main(){
	result := Add("ron","jon")
	fmt.Println(result)
}