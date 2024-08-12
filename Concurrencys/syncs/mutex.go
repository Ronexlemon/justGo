package syncs

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

var count int=0

func increment(){
	//acquire the lock
	mutex.Lock()
	defer mutex.Unlock()
	
	count ++
	count ++
	
	fmt.Println("the new count is before ",count)

}

func decrement(){
	//acquire the lock
	mutex.Lock()
	defer mutex.Unlock()
	count --
	fmt.Println("the new count after dec is",count)

}

func Mutex(){
	var wg sync.WaitGroup

	wg.Add(2)

	go func(){
		increment()
		defer wg.Done()

	}()
	go func(){
		decrement()
		defer wg.Done()

	}()

	wg.Wait()
	fmt.Println("final count is",count)

}


func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d starting\n",id)

	time.Sleep(time.Second)

	fmt.Printf("Worker%d done\n",id)
}

func Mutex2(){
	var wg sync.WaitGroup

	for i:=0;i<=5;i++{
		wg.Add(1)
		go worker(i,&wg)
	}
	wg.Wait()
	fmt.Println("All workers finished execution")

}