package channel03

import (
	"fmt"
	"sync"
	"time"
)
//dowork

func doWork(){
	time.Sleep(1*time.Second)
}
//pass without
func fetchDataWithout(cities string)string{
	doWork()
	return  cities
}

func WithoutRoutineAndChannel(){
	startTime :=time.Now()
	cities :=[]string{"japan","Nairobi","tokyo"}

	for _,city:=range cities{
		data:= fetchDataWithout(city)
		fmt.Println(data)
	}
	fmt.Println("time taken",time.Since(startTime))

}

//pass with channels
func fetchData(cities string,c chan<-string,wg *sync.WaitGroup){
	defer wg.Done()
	doWork()

	
		c <- fmt.Sprintf("the value is %s",cities)
	
}

func RoutinesChannel(){
	startTime :=time.Now()
	c := make(chan string)

	cities :=[]string{"japan","Nairobi","tokyo"}
var wg sync.WaitGroup
	for _,city:=range cities{
		wg.Add(1)
		go fetchData(city,c,&wg)
	}
	go func ()  {
		wg.Wait()
		close(c)
	}()

	for i:=range c{
		fmt.Println(i)
	}
	fmt.Println("time taken",time.Since(startTime))
}