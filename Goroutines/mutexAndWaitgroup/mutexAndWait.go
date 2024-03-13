package mutexandwaitgroup

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}
var mutex = &sync.Mutex{}
func MutexAndWait(){
	
	fmt.Println("Race condition - ------------------------ lemonr")
	score:=[]int{0}

	wg.Add(3)
  go	func(wg *sync.WaitGroup,m *sync.Mutex) {
	fmt.Println("One R")
	m.Lock()
	score = append(score, 1)
	m.Unlock()
	wg.Done()
  }(wg,mutex) 
  
  go	func(wg  *sync.WaitGroup,m *sync.Mutex){
	fmt.Println("Two R")
	m.Lock()

	score = append(score, 2)
	m.Unlock()
	wg.Done()
  }(wg,mutex) 
  go	func(wg *sync.WaitGroup,m *sync.Mutex){
	fmt.Println("Three R")
	m.Lock()
	score = append(score, 3)
	m.Unlock()
	wg.Done()
  }(wg,mutex) 

  wg.Wait()
  fmt.Println(score) // [1 2 3]

}