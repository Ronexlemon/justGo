package syncs

import (
	"fmt"
	"sync"
)


type safeCounter struct{
	mu sync.Mutex
	count int
}

//method to right

func(s *safeCounter) increment(){
	s.mu.Lock()
	
	s.count ++
	s.mu.Unlock()
}

func(s *safeCounter) decrement(){
	s.mu.Lock()
	
	s.count --
	s.mu.Unlock()
}

func(s *safeCounter) readValue()int{
	s.mu.Lock()
	
	defer s.mu.Unlock()
	return s.count
}

func MutexAndWaitGroup(){
	var wg sync.WaitGroup
	counter := safeCounter{}

	for i:=0;i< 1000;i++{
		wg.Add(2)
		go func(){
			counter.increment()
			defer wg.Done()

		}()
		go func(){
			counter.decrement()
			defer wg.Done()

			}()
	}
	wg.Wait()

	fmt.Println("final counter is",counter.readValue())
}