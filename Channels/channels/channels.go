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

func BufferedChannels(){

	dataChan:= make(chan A,2)

	dataChan <- A{678,"john doe"}
	dataChan <- A{678,"john de"}

	results := <- dataChan
	result2 := <- dataChan
	fmt.Println(results)
	fmt.Println(result2)
}