package main

import (
	"gomethods-interfaces/interfaces"
	"gomethods-interfaces/methods"
	"gomethods-interfaces/structs"
)

func main(){
	structs.GetCircleDetails()
	methods.MethodCircleDetails()
	methods.MethodAnotherStruct()
	interfaces.GetLogs()
	interfaces.Receipt()
}