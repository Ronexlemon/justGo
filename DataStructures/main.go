package main

import (
	"fmt"
	"godatastructures/arrays"
	"godatastructures/slices"
)


func main(){
	arrays.Array()
	arrays.FixedArray()
	arrays.MultiDimension()
	arrays.DetermineLen()
	slices.InitSlice()
	slices.InitSlice2()
	slices.AppendTo()
	slices.Copy()
	fmt.Println(slices.Compareslices(slices.Slice,slices.Slice2))
}

