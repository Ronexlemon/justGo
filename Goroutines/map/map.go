package mapping

import (
	"fmt"

	
)

func Mapping(){

	fmt.Println("Mapps ing go")

	languages := make(map[string]string) //Declare 
	languages["math"] = "mathmatics"
	languages["py"] = "python"
	languages["sol"] = "solidity"

	fmt.Println("Lit of all languages",languages)
	fmt.Println("value of List of all languages",languages["sol"])

	delete(languages,"math")// delete a key value
	fmt.Println("Lit of all languages",languages)

	//loops
	for key,value:=range languages{
		fmt.Printf("%s is %s\n",key,value)
	}



	
	
}



func newFunc(){
	fmt.Println("new function")
}

