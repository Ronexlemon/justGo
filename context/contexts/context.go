package contexts

import (
	"context"
	"fmt"
	"sync"
	"time"
)
func func1(ctx context.Context,parentWait *sync.WaitGroup,stream <-chan interface{}){
	defer parentWait.Done()
	var wg sync.WaitGroup

	doWork := func(ctx context.Context){
		defer wg.Done()
		for {
			select{
			case <- ctx.Done():
				return
			case d,ok:= <-stream:
				if !ok{
					fmt.Println("Channel closed")
					return
				}
				fmt.Println(d)
			}
		}
	}

	newCtx,cancel := context.WithTimeout(ctx,time.Second*3)
	
	defer cancel()
	for i :=0;i <3;i++{
		wg.Add(1)
		doWork(newCtx)
	}
	wg.Wait()

}

func genericFunc(ctx context.Context,wg *sync.WaitGroup,stream  <-chan interface{}){
	defer wg.Done()
	for {
		select{
			case <- ctx.Done():
				return
			case d,ok := <-stream:
				if !ok{
					fmt.Println("Channel closed")
					return
					}
					fmt.Println(d)
		}
	}
}
func Context(){
	var wg sync.WaitGroup

	ctx,cancel:= context.WithCancel(context.Background())

	defer cancel()

	generator := func(dataItem string, stream chan interface{}){

		for{
			select{
			case <- ctx.Done():
				return
			case stream <-dataItem:

			}
		}
	}

	inifineApple := make(chan interface{})

	go generator("Apples",inifineApple)
	inifineoranges := make(chan interface{})

	go generator("Oranges",inifineoranges)
	inifineMango := make(chan interface{})

	go generator("Mango",inifineMango)
	wg.Add(1)

	go func1(ctx,&wg,inifineApple)
	func2 := genericFunc
	func3:= genericFunc

	wg.Add(1)
	go func2(ctx,&wg,inifineMango)
	wg.Add(1)
	go func3(ctx,&wg,inifineoranges)
	wg.Wait()



}

