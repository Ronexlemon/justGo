package tests

import "fmt"

//write a program that finds the smallest number in this list

var XX =[]int{48,96,86,68,57,82,63,70,28,37,34,83,79,27,19,97,9,7}


func findSmallestNumber(list []int)int{
	if len(list) == 0 {
		return 0 // or handle the empty slice case as needed
	}
	smallNum := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] < smallNum {
			smallNum = list[i]
		}
	}
	return smallNum
	

}

func FindSmallest(){
	small:= findSmallestNumber(XX)
	fmt.Println("smallest number is",small)
}