package channel

import (
	"fmt"
	"time"
)


func SelectChannel(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		for{
			ch1 <- "from 1"
			time.Sleep(time.Second*2)
		}

	}()
	go func(){
		for{
			ch2 <- "from 2"
			time.Sleep(time.Second*3)
		}

	}()

	go func(){
		for{
			select{
			case mg1 :=<-ch1:
				fmt.Println(mg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
			}
		}

	}()
	var input string
	fmt.Scanln(&input)
}