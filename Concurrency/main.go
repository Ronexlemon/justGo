package main

import (
	"fmt"
	"time"
)

func main(){
	welcome:="Welcome to go concurrency"
	fmt.Println(welcome)
	go greeter("Hello")
	 greeter("Yollow")
}

func greeter(s string){
	for i:=0; i< 6;i++{
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}