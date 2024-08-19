package main

import (
	"fmt"
	"gopackages/greeting"
)


func main(){
	var a int = 10
	var b int = 20

	num,err:=greeting.SubTraction(a,b)

	if err !=nil{
		fmt.Println(err.Error())
		//panic(err) wiil exits
	}
	fmt.Println(num)

}