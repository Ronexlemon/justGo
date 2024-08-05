package switchcontrol

import "fmt"


//select multiplication of a number upto  10

func Switch(num int){
	// Switch statement
	switch num{
	case 0: NumberMultiplication(0)
	case 1: NumberMultiplication(1)
	case 2: NumberMultiplication(2)
	case 3: NumberMultiplication(3)
	case 4: NumberMultiplication(4)
	case 5: NumberMultiplication(5)
	case 6: NumberMultiplication(6)
	case 7: NumberMultiplication(7)
	default:NumberMultiplication(num)
	}
}


func NumberMultiplication(num int){


	for i:=0; i<num; i ++{
		fmt.Println(num,"*",i,"=",num*i)
	}

}