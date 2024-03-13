package main

import (
	"fmt"
	"sync"

	"net/http"
	//"time"
)

// func greeter(s string){
// 	for i:=0;i<6;i++{
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }
//mutex
var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Opps we got an error")
	}
	mut.Lock() //lock the resource to make it thread
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}
func main() {
	// go greeter("yollow")
	//  greeter("yes")
	urls := []string{"https://google.com", "https://facebook.com", "https://safaricom.com", "https://github.com/RonexLemon"}

	for _, web := range urls {
		 go getStatusCode(web)
		 wg.Add(1)
	}
wg.Wait()
fmt.Println(signals)
}
