package controlflows

import "fmt"
type Types struct{
	Name string
	Age int
	

}

func Loops(){
	welcome:= "loops is looping"
	fmt.Println(welcome)

	days := []string{"sunday","Tuesday","wednesday","friday","saturday"}
	fmt.Println(days)

	// ////forloop
	// for d:=0; d< len(days);d++{
	// 	fmt.Println("days ", days[d])
	// }
	//for range
	for key,value := range days{
		fmt.Printf("the key is %d and value %s\n",key,value)
	}

	//while for
	value:=1
	value2:= 3

	for value <10{
		if value ==6{
			value ++
			continue
		}
		
		fmt.Println("value is",value)
		value ++
	}
}