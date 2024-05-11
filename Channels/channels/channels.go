package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type A struct {
	age  int
	name string
}

var wg sync.WaitGroup

func Channels() {
	dataChan := make(chan A)

	go func() {
		dataChan <- A{789, "yollow"}

	}()

	result := <-dataChan

	fmt.Println(result)

}

func BufferedChannels() {

	dataChan := make(chan A, 2)

	dataChan <- A{678, "john doe"}
	dataChan <- A{678, "john de"}

	results := <-dataChan
	result2 := <-dataChan
	fmt.Println(results)
	fmt.Println(result2)
}

func ChannelsSimultenous() {
	dataChan := make(chan A)

	go func() {
		dataChan <- A{789, "yollow"}
		dataChan <- A{75, "yollow"}
		dataChan <- A{73, "yollotw"}
		close(dataChan)

	}()

	for n := range dataChan {

		fmt.Println(n)
	}

	result := <-dataChan

	fmt.Println(result)

}

func DoWork() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func ChannelWithLoad() {
	datachan := make(chan A)

	go func(wg *sync.WaitGroup) {
		
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				datachan <- A{result, "yollow"}

			}()
		}
		wg.Wait()
		close(datachan)
	}(&wg)

	for n := range datachan {
		fmt.Println(n)
	}

}
