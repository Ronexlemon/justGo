package greeting

import "errors"


func SubTraction(num int, num2 int)(int,error){
	//the first number should be greater than the second

	if num < num2{
		return 0,errors.New("the first number should be greater than the second")
	}
	return (num - num2), nil


}