package channel

import (
	"fmt"
	"time"
)


func clientA(r chan string, done chan bool){
	fmt.Println("Client A Typing ....")
var input string
fmt.Scanln(&input)
time.Sleep(time.Millisecond*250)
r<-input
if input == "done" {
	done <- true
}
	
}

func clientB(r chan string,done chan bool){
	fmt.Println("Client B Typing ....")
	var input string
	fmt.Scanln(&input)
	time.Sleep(time.Millisecond*250)
	r<-input
	if input == "done" {
		done <- true
	}
		
	}


	func Chat(){
		r := make(chan string)
		done := make(chan bool,1)

		

		go func(){
			for{
				select{
					case msg := <-r:
						fmt.Println(msg)
						
					
				}
			}
		}()
		for {
			select {
			case <-done:
				fmt.Println("Closing Chat")
				close(r)
				return
			default:
				clientA(r, done)
				clientB(r, done)
			}
		}

		

	}