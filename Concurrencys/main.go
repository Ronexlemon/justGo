package main

import (
	"goconcurrency/channel"
	"goconcurrency/goroutines"
)

func main(){
	goroutines.Goroutine()
	goroutines.MultipleGoroutines()
	channel.Channel()
}