package greeting

import (
	"testing"
)

func Test_If_FirstNumber_is_Not_Greater_Than_the_Second(t *testing.T) {
	var a int = 10
	var b int = 20

	result, err := SubTraction(a, b)
	if err == nil {
		t.Fatalf("Expected an error, but got nil")
	}

	expectedError := ErroName
	if err != expectedError {
		t.Fatalf("Expected error %v, but got %v", expectedError, err)
	}

	// Check that the result is zero when there is an error
	if result != 0 {
		t.Fatalf("Expected result 0, but got %d", result)
	}
}


//test for the correct values

func Test_For_The_Correct_Subtraction(t *testing.T){
	var a int = 20
	var b int = 10

result,err := SubTraction(a,b)
if err !=nil{
	t.Fatalf("Expected no error, but got %v",err)
}
expectedResult := a-b
if result != expectedResult{
	t.Fatalf("Expected result %d, but got %d",expectedResult,result)
	}
}