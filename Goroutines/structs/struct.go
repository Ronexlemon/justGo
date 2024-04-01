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
	//fmt.Println(lemon.CreateUser("john",20,"doe@gmail.com"))
	//accessing the fields of a struct
	fmt.Println(lemon.Name)
	fmt.Println(lemon.Age)

}

func (u *User) CreateUser(name string,age int,email string)User{
	u.Age = age
	u.Name = name
	u.Email= email
	return *u


}
func (u User) GetStatus(){
	fmt.Println("is userActive",u.Status)

}

func (u User) NewMail(){
	u.Email="newmail@gmail.com"
	fmt.Println("New Mail is ", u.Email)
}
