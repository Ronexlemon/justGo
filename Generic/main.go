package main

import (
	"fmt"
	"gogeneric/mapping"

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
	
	fmt.Println(mapping.MapValues([]int{1,2,3,4,5},func(n int)int{return n*2}))
	fmt.Println(result)
}