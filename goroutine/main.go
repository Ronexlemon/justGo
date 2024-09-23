package main

import (
	"fmt"
	goroutine "goroute/goroute"
	"goroute/pattern"
)

func main(){
	fmt.Println("goroutines")
	goroutine.GoRoute()
	goroutine.Chan()
	goroutine.Select()
	pattern.ForSelectLoops()
	pattern.DoneChannel()
	pattern.Pipeline()

	
}