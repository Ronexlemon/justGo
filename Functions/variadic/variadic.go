package variadic

import "fmt"

//A Special form available for the last parameter in a Go function
//  By using ... before the type name of the last parameter you can inidicate that it takes zero or more
// of those parameter

func Add(args ...int){

	sum:=0

	for _,value:= range args{
		sum+=value
	}

	fmt.Println("the sum is",sum)
}