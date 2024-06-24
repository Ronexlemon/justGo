package interfaces

import (
	"math/rand"
	"fmt"
)

type Player interface{
	KickBal()
}

type FootbalPlayer struct{
	stamina int
	 power int
}

type Messi struct{
	FootbalPlayer
	Speed int

}
func (f Messi)KickBal(){
	shot := f.stamina +f.power * f.Speed

	fmt.Println(" Messi The Player Shot is: %\v", shot)
}

func (f FootbalPlayer)KickBal(){
	shot := f.stamina +f.power

	fmt.Println("The Player Shot is:", shot)
}

func Interfaces(){

	team:= make([]Player,11)

	for i:=0; i <len(team)-1; i++{
		team[i] = FootbalPlayer{
			stamina:  rand.Intn(10),
			power: rand.Intn(10),
		}
	}
	
	
	team[len(team)-1] = Messi{
		FootbalPlayer: FootbalPlayer{
			stamina: 10,
			power: 10,
		},
		Speed: 10,
	}
	for _,player :=range team{
		player.KickBal()
	}
}