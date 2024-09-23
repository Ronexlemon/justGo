package pattern

import (
	"fmt"
	"time"
)

// For-select loop
// Done Channel
// Pipelines
var Char []string = []string{"a", "b", "c", "d"}
var CharB []string = []string{"a", "b", "c", "d"}

func ForSelectLoops() {
	//For-select loop
	//buffered channel
	charChannel := make(chan string, len(Char))
	for _, s := range Char {
		select {
		case charChannel <- s:
		}

	}
	close(charChannel)
	for result := range charChannel {
		fmt.Println(result)
	}

	for result2 := range charChannel {
		fmt.Println(result2)
	}

}

//done Channel

func doWork(done <-chan bool) {
	select {
	case <-done:
		return
	default:
		fmt.Println("doing Work")
	}
}

func DoneChannel() {
	done := make(chan bool)
	go doWork(done)

	defer close(done)
	time.Sleep(time.Second * 10)
}

//PipeLine

func sliceToChannel(nums []int)<-chan int{
	out:= make(chan int)
	go func() {
		for _,value := range nums{
			out<-value
		}
		close(out)
	}()
	return out

}
func sq(nums <-chan int)<-chan int{
	square:=make(chan int)
	go func ()  {
		for n:= range nums{
			square<-n*n
		}
		close(square)
		
		
	}()
	return square
	
}
func Pipeline(){
	//input
	nums := []int{1,2,3,4,5}
	//Stage 1

	dataChannel :=  sliceToChannel(nums)

	//Stage 2
	finalchannel:= sq(dataChannel)
	//stage 3
	for n:= range finalchannel{
		fmt.Println(n)
	}
}