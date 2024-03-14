package channels

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func Channel(){
	fmt.Println("Channels in golang ------lemonr")
	myChan := make(chan int,5)

	wg.Add(2)
	//Receive only

	go func(ch <-chan int,wg *sync.WaitGroup){ // receiving a value
		val,isChannelopen := <-myChan  //listening to channel
		fmt.Println(isChannelopen)
		fmt.Println(val)
		             //fmt.Println(<-myChan)

   
		wg.Done()

	}(myChan,wg)

	//Send-Only
	go func(ch chan<- int,wg *sync.WaitGroup){ //sending a value
	     myChan <- 0
		// myChan <- 6
		// myChan <- 5
		// myChan <- 6

        close(myChan)
		wg.Done()

	}(myChan,wg)

	// myChan <- 5   //push 5 to myChan

	// fmt.Println(<-myChan) //receive value  from myChan, this will print 5 this will block  if there is no data available on the channel


	wg.Wait()
}