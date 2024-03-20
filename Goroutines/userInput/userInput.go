package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput(){
	Welcome:= "Welcome to user input"
	fmt.Println(Welcome)
	reader :=  bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for celo:")

	//comma ok || err err
	input,err:= reader.ReadString('\n') //reads until it encounters a newline character '\n'
	if err !=nil{
		panic(err)
	}
	fmt.Println("thank for reading", input)

}