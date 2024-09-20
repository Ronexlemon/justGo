package goroutine

import (
	"fmt"
	"time"
)

func doSomething(num string){
	fmt.Println(num)
}

func GoRoute(){
	go doSomething("Hello One")
	go doSomething("Hello Two")
	go doSomething("Hello Three")
	time.Sleep(time.Second *2)
	fmt.Println("Done .....")
}