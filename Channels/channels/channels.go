package channels

import "fmt"


type A struct{
   age	 int 
	name string
	
}



func Channels(){
	dataChan := make(chan A)

	

	go func() {
		dataChan <- A{789,"yollow"}

	}()
	

	result := <- dataChan


	fmt.Println(result)

}