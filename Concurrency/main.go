package main

import (
	"fmt"
	"goconcurrency/racecondition"
	"net/http"
	"sync"
	// "time"
)

var wg sync.WaitGroup //pointer
var mut sync.Mutex  //pointer
var signals = []string{"test"}

func main(){
	welcome:="Welcome to go concurrency"
	fmt.Println(welcome)

	
	racecondition.RaceCondition()
	
	
}

// func greeter(s string){
// 	for i:=0; i< 6;i++{
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func callStatusCode(){
	websiteList := []string{
		"https://www.google.com",
		"https://www.github.com",

		"https://www.stackoverflow.com",
		"https://www.udemy.com",
		"https://www.golangprograms.com",
		
	}
	for _,value:=range websiteList{
		go getStatusCode(value)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println("signals",signals)
}

func getStatusCode(endpoint string){
	
	defer wg.Done()
	res,err:=http.Get(endpoint)
	checkNill(err)
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d status code for %s \n",res.StatusCode,endpoint)
}

func checkNill(err error){
	if err!=nil{
		panic(err)
	}
}