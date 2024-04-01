package structs

import "fmt"

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func Structs() {
	welcome := "Welcome to structs"
	fmt.Println(welcome)

	lemon := User{Name: "ronex",Age: 23,Status: true,Email: "ronexlemon@gmail.com"}
	

	fmt.Println(lemon)
	fmt.Printf("my name is %v  and i am %d years old already verified %v and my email %s",lemon.Name,lemon.Age,lemon.Status,lemon.Email)
	fmt.Println(lemon.createUser("john",20,"doe@gmail.com"))
	//accessing the fields of a struct
	fmt.Println(lemon.Name)
	fmt.Println(lemon.Age)

}

func (u *User) createUser(name string,age int,email string)User{
	u.Age = age
	u.Name = name
	u.Email= email
	return *u


}
