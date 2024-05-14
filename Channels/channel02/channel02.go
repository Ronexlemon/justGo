package channel02

import "fmt"

//ch <- v -> send v to channel ch
// v:= <- ch -> receive fromch, and assign value to v
//NB -> The data flow in the direction of the arrow
//like map and slices channels must be created before use


func sum(s []int, c chan int){
	sum:=0

	for _,value:=range s{
		sum+=value
	}
	c <-sum
}

func Channel02(){

	s :=[]int{1,2,3,4,5,6,-6}

	c := make(chan int)

	go sum(s,c)
	go sum(s,c)

	x,y := <-c,<-c

	fmt.Println(x,y,x+y)
}

func BuffereChan(){
	ch := make(chan int,2)
	ch <- 1 
	ch <- 2
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
}