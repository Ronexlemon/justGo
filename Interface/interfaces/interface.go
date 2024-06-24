package interfaces

import (
	"math/rand"
	"fmt"
)


type FootbalPlayer struct{
	stamina int
	 power int
}


func (f FootbalPlayer)KickBal(){
	shot := f.stamina +f.power

	fmt.Println("The Player Shot is:", shot)
}

func Interfaces(){

	team:= make([]FootbalPlayer,11)

	for i:=0; i <len(team); i++{
		team[i] = FootbalPlayer{
			stamina:  rand.Intn(10),
			power: rand.Intn(10),
		}
	}
	for _,player :=range team{
		player.KickBal()
	}
}