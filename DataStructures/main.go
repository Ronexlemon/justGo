package main

import (
	"fmt"
	"godatastructures/arrays"
	"godatastructures/maps"
	"godatastructures/slices"
	"godatastructures/tests"
)


func main(){
	fmt.Println("***************A*R*R*A*Y*S")
	arrays.Array()
	arrays.FixedArray()
	arrays.MultiDimension()
	arrays.DetermineLen()
	fmt.Println("***************S*L*I*C*E*S")
	slices.InitSlice()
	slices.InitSlice2()
	slices.AppendTo()
	slices.Copy()
	fmt.Println(slices.Compareslices(slices.Slice,slices.Slice2))
	fmt.Println("***************M*A*P")
	maps.MapInit()
	maps.GetKeysValues(maps.ReturnAMap())
	maps.DeleteMapAtAnIndex(maps.ReturnAMap(),"age")

	tests.FindSmallest()

}

