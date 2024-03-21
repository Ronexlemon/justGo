package timee

import (
	"fmt"
	"time"
)

func  Time(){
	welcome:= "Welcome to the Time Elapsed"
	fmt.Println(welcome)
	presentTime:= time.Now()

	fmt.Println(presentTime)

	//format time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //always the format

	
}