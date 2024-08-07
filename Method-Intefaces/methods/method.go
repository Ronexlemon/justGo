package methods

import (
	"fmt"
	"math"
)

type Circle struct{
	rad float64
}

type Person struct{
	Name string
}
type Android  struct{
	Person
	Phone string
}

func (c *Circle) area()float64{
	return  math.Pi * c.rad *c.rad
}


func (p *Person) Talk(){
	fmt.Println("Hello, my name is ", p.Name)
}


func MethodCircleDetails(){
	c := &Circle{rad: 30}
	fmt.Println("New Area",c.area())
}
func MethodAnotherStruct(){
	a := new(Android)
	a.Name = "john doe"
	a.Phone = "Iphone 15"
	a.Person.Talk()

}