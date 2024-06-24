package interfaces

import (
	"math/rand"
	"fmt"
)

type Player interface{
	KickBal() int
	Name() string
}

type FootbalPlayer struct{
	stamina int
	 power int
}

type Messi struct{
	FootbalPlayer
	Speed int

}
func (m *Messi)Name()string{return "Messi"}

func (f *FootbalPlayer)Name()string{ return "Random"}
func (f *Messi)KickBal()int{
	shot := f.stamina +f.power * f.Speed

	return shot

	
}

func (f *FootbalPlayer)KickBal()int{
	shot := f.stamina +f.power

	return shot
}

func Interfaces(){

	team:= make([]Player,11)

	for i:=0; i <len(team)-1; i++{
		team[i] = &FootbalPlayer{
			stamina:  rand.Intn(10),
			power: rand.Intn(10),
		}
	}
	
	
	team[len(team)-1] = &Messi{
		FootbalPlayer: FootbalPlayer{
			stamina: 10,
			power: 10,
		},
		Speed: 10,
	}
	for _,player :=range team{
		
		fmt.Printf("%s is kicking the ball %d\n", player.Name(),player.KickBal())
	}
}