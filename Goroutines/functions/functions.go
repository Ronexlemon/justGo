package functions

import "fmt"

func Functions(){
	welcome:="welcome to functions yollow!"
	fmt.Println(welcome)
	fmt.Println(addNum(3,5))
	fmt.Println(proAdder(1,5,7,9,2,5,7,8))

	

}

func  addNum(num1 int, num2 int)int{
	return num1 +num2
}

func  addNum2(num1 int, num2 int)int{
	return num1 +num2
}

func proAdder(values ...int)int{
	total:= 0

	for _,value:= range values{
		total +=value
	}

	return total
}