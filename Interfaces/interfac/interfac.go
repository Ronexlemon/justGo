package interfac

import (
	"fmt"
	"math"
)

type tank interface {
	TankArea() float64
	Volume() float64
}

type myValue struct {
	radius float64
	height float64
}

// implement methods of
// the tank interface
func (v myValue) TankArea() float64 {
	return 2*v.radius*v.height + 2*math.Pi*v.radius*v.radius

}

func(v myValue) Volume()float64{
	return math.Pi*v.radius*v.radius*v.height

}

func passInterface(a interface{}){
	val:=a.(string)
	fmt.Println(val)

}
//assertion switch

func assertionSwitch(a interface{}){
	switch a.(type){
	case int:
		fmt.Println("int",a.(int))
	case string:
		fmt.Println("string",a.(string))
	case float64:
		fmt.Println("float",a.(float64))
	default:
		fmt.Println("default")

	}
}

func Int(){
	var t  tank
	t =myValue{10,20}
	fmt.Println("Tank area:",t.TankArea())
	fmt.Println("Tank volume:",t.Volume())
}
func InterfaceAssertion(){
	var v1 interface{} = "yes lemonr"
	var v2 interface{} = 909.90
	passInterface(v1)
	assertionSwitch(v2)
}