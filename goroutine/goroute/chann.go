package goroutine

import "fmt"



func Chan(){
	myChan := make(chan string)
	defer close(myChan)

	go func(){
		myChan <- "Yes Yes Yes Receive some "
	}()
	msg := <-myChan

	fmt.Println(msg)

}