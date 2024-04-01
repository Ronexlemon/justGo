package methods

import ( "fmt"

"goroutines/structs")




func Methods(){
	welcome:= "welcome to methods yollow!"
	fmt.Println(welcome)
	
	john:= structs.User{Name: "john doe",Age: 20,Status: true,Email: "does@gmail.com"}
	john.GetStatus()
	john.NewMail()
	fmt.Printf("my name is %v  and i am %d years old already verified %v and my email %s",john.Name,john.Age,john.Status,john.Email)
	
	
	
	
}

