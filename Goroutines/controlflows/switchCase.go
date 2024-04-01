package controlflows

import (
	"fmt"
	"math/rand"
	"time"
)

func SwitchCase(){

	welcome:= "switch case yollow!!"
	fmt.Println(welcome)
	//rand.Seed(time.Now().UnixNano()) //deprecated
	rand.NewSource(time.Now().UnixNano()) //

	diceNumber:= rand.Intn(6)+1
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber{
	case 1: 
	fmt.Println(welcome+ "move one step and you can open")
	
	case 2 : fmt.Println(welcome +"move two steps")
	case 3: fmt.Println(welcome+"move three steps")
	case 4 : fmt.Println(welcome+" and move four steps too.")
	case 5: fmt.Println(welcome+ " and move five steps too.")
	case 6: fmt.Println(welcome + "move six steps and play again")
	default: fmt.Println("you are not allowed to play the game.")
	}
}