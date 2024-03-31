package slices

import (
	"fmt"
	"sort"
)

func Slices(){
	fmt.Println("Welcome to slices")

	var fruitList =[]string{"Apple","tomato"}
	fruitList=append(fruitList,  "banana", "Mango")
	fmt.Printf("type of fruitlist is %T\n", fruitList)
	fmt.Println("fruitlist slice:",fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int,4)

	highScores[0] = 234
	highScores[1] = 423
	highScores[2] = 945
	highScores[3] = 342
	//highScores[4] = 777
	highScores = append(highScores, 555,666)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) 

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // returns true if sorted

	//remove value based on index

	var courses  = []string{"reactjs","javascript","python","ruby"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...) //skip that index
	fmt.Println(courses)
	


}