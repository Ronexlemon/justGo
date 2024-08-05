package main

import (
	"gocontrol/forcontrol"
	"gocontrol/ifelse"
	"gocontrol/switchcontrol"
)


func main(){
	ifelse.IfElse()
	forcontrol.ListNumber()
	forcontrol.MuliplicationTable()
	ifelse.ListNumbersIfEvenOrOdd()
	switchcontrol.Switch(5)
	switchcontrol.Switch(50)

}