package main

import (
	//"goconcurrency/channel"
	"goconcurrency/syncs"
	//"goconcurrency/goroutines"
)

func main(){
	// goroutines.Goroutine()
	// goroutines.MultipleGoroutines()
	// channel.Channel()
	// channel.SelectChannel()
	//channel.Synchronize()
	//channel.Chat()
	syncs.Mutex()
	syncs.Mutex2()
	syncs.MutexAndWaitGroup()
}