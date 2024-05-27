package main
/**
*****Non-Pointer Values ****                                                                   Pointer wrapper  values
1.    Strings                                                                     1.   slices
2.    ints                                                                        2.    maps
3.    boolean                                                                       3.    functions
4.    floats                                                                      
5.                                                                        
6.    arrays                                                                       

****/

import (
	"fmt"
)
func updateName(name *string){
	*name = "Rahul"
}

func main() {
	welcome := "Go pointers"
	fmt.Println(welcome)
	name:= "lemonr"
	
	// fmt.Println("memory address of name is",&name)
	m:=&name
	
	fmt.Println("memory address of m is",m)
	fmt.Println("value  of name is",*m) //

	updateName(m)
	fmt.Println(name)


	
}
