package structs

import (
	"fmt"
	"math"
)

type Circle struct{
	radius float64
}

//area of a circle

func area(c *Circle)float64{
	return math.Pi * c.radius *c.radius
}

func perimeter(c *Circle)float64{
	
	return 2 *math.Pi *c.radius
}

func GetCircleDetails(){
	c:= Circle{radius: 10.00}
	ar:= area(&c)
	sum:= perimeter(&c)

	fmt.Println("the areas is",ar)
	fmt.Println("the perimeter is",sum)
	fmt.Println("the radies is",c.radius)

}