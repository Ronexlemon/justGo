package racecondition

import (
	"fmt"
	"sync"
)

func RaceCondition() {
	fmt.Println("Race condtion")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup,mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

	}(wg,mut)
	go func(wg *sync.WaitGroup,mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

	}(wg,mut)
	go func(wg *sync.WaitGroup,mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

	}(wg,mut)

	wg.Wait()
	fmt.Println(score)

}
