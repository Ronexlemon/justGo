package pattern

import "fmt"

//For-select loop
//Done Channel
//Pipelines
var Char []string = []string{"a","b","c","d"}
var CharB []string = []string{"a","b","c","d"}


func ForSelectLoops(){
	//For-select loop
	//buffered channel
	charChannel := make(chan string,len(Char))
	for _,s := range Char{
		select{
			case charChannel <- s:
		}
		
	}
	close(charChannel)
	for result := range charChannel{
		fmt.Println(result)
	}
	 
}

//done Channel