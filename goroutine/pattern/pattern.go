package pattern

import (
	"fmt"
	"math/rand"
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

//CHAPTER 2

func Game(){
	start:= time.Now()

	defer func ()  {
		fmt.Println(time.Since(start))
		
	}()
	players := []string{"Lemonr","John Doe","Kales","Ban"}
	for _,player:= range players{
		go attack(player)
	}
time.Sleep(time.Second *1)
}
func Gamev1(){
	start:= time.Now()
	defer func ()  {
		fmt.Println(time.Since(start))}()

		smokeItOut := make(chan bool)
		monster := "Hakes"
		go attackv1(smokeItOut,monster)
	fmt.Print(<-smokeItOut)
	close(smokeItOut)
}

func attackv1(done chan bool,monster string){
	time.Sleep(time.Second) // prepare for launch
	fmt.Println(monster,"is attacked")
	done <- true
	

}
func attack(player string){
	fmt.Println(player,"is attacking")
	time.Sleep(time.Second)
	
}

//channelV2 -> buffer
func Gamev2(){
	channel := make(chan string,2)
	channel <- "First"
	channel <- "Second"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	
}
func Gamev2_1(){
	channel := make(chan string)
	numRounds := 3
	go throwDart(channel,numRounds)
	for i:=0; i< numRounds ;i++{
		fmt.Println(<-channel)
	}
	

}

func throwDart(channel chan string, rounds int){
	rand.Seed(time.Now().UnixNano())
	
	for i:=0; i< rounds ;i++{
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored",score)

	}
	

}