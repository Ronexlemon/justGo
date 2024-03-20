package variables

import "fmt"

const  LoginToken string = "ghghghghght" //public

func Variables(){
	var isLogging bool = true
	fmt.Println(isLogging)
	fmt.Printf("This varibale is of type %T \n", isLogging)


	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("This varibale is of type %T \n",smallValue)

	var smallFloat float64 = 255.569697977977
	fmt.Println(smallFloat)
	fmt.Printf("This varibale is of type %T \n",smallFloat)

	//default value
	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("This varibale is of type %T \n",anotherValue)

	//implicit way to init a variable

	var  name = "john doe"
	fmt.Println(name)

	//no var style

	numberOfUsers:= 4040
	fmt.Println(numberOfUsers)
	

}