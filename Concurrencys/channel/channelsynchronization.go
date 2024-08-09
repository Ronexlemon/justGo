package channel

import (
	"fmt"
	"time"
)


func WaitForDone(done chan <-bool){
	fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
	done<- true
} 

func Synchronize(){
	done:= make(chan bool,1)
	go WaitForDone(done)

	<-done



}