package array

import "fmt"

func Arrays(){
	fmt.Println("Welcome to array");

	var fruits [4]string

	fruits[0]= "mango"
	fruits[1] = "tomatio"
	fruits[3] = "orange"


	fmt.Println("fruit list is:", fruits)
	fmt.Println("fruit list is:", len(fruits))

	var vegList = [3]string{"potato","beans","kales"}

	fmt.Println("vegList list is:", vegList)
	fmt.Println("vegList list is:", len(vegList))
}