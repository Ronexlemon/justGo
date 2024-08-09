package channel

import (
	"fmt"
	"time"
)

//channel provide a way for two goroutines to communicate with one another and synchronize their execution


func pinger(c chan <-string){
	for i:=0;;i++{
		c <- "ping"
	}
}


func printer(c <-chan  string){
	for {
		msg:= <-c
		fmt.Println(msg)
		time.Sleep(time.Second*1)
	}
}

func pong(c chan <-string){
	for i:=0;;i++{
		c <- "pong"
	}
}

func Channel(){
	var c chan string = make(chan string)

	go pinger(c)
	go pong(c)
	go printer(c)
	// var input string
	// fmt.Scanln(&input)

	time.Sleep(time.Second*10)
}

