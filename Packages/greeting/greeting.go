package greeting

import (
	
	"fmt"
)

var ErroName  = fmt.Errorf("the first number should be greater than the second")
func SubTraction(num int, num2 int)(int,error){
	//the first number should be greater than the second

	if num < num2{
		return 0,ErroName
	}
	return (num - num2), nil


}