package goroutines

import "fmt"




func f(i int){
	for g:=0;g <i; g++{
		fmt.Println("all numbers",g)
	}
}


func Goroutine(){
	go f(20)
	
	var input string
	 fmt.Scanln(&input)
	fmt.Println("the value inputed is",input)
	
	fmt.Println("done")
}

