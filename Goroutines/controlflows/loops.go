package controlflows

import "fmt"

func Loops(){
	welcome:= "loops is looping"
	fmt.Println(welcome)

	days := []string{"sunday","Tuesday","wednesday","friday","saturday"}
	fmt.Println(days)

	////forloop
	for d:=0; d< len(days);d++{
		fmt.Println("days ", days[d])
	}
	
}