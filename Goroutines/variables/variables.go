package variables

import "fmt"

func Variables(){
	var isLogging bool = true
	fmt.Println(isLogging)
	fmt.Printf("This varibale is of type %T", isLogging)
}