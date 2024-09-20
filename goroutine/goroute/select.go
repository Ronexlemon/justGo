package goroutine

import "fmt"



func Select(){
	myChan := make(chan string)
	anotherChan := make(chan string)

	go func(){
		myChan <- "Yes data from MyChan"
	}()
	go func(){
		anotherChan <- "Yes data from AnotherChan"
		}()
	select{
	case datafromMyChannel :=  <-myChan:
		fmt.Println(datafromMyChannel)
	case datafromAnotherChannel := <-anotherChan:
		fmt.Println(datafromAnotherChannel)
	}

}