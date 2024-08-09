package goroutines

import (
	"fmt"
	//"math/rand"
	"time"
)




func f(i int){
	for g:=0;g <i; g++{
		fmt.Println("all numbers",g)
		
	}
}


func Goroutine(){
	go f(20)
	
	
	
	fmt.Println("done")
	 time.Sleep(time.Millisecond*10)
}

func MultipleGoroutines(){
	fmt.Println("New Routines")
	for i:=0;i<10;i++{
		go f(10)
	}
	time.Sleep(time.Millisecond*10)
}