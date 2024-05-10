package structs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)


type CustomData interface{
	constraints.Ordered |[]byte | []rune
}

type User[T CustomData] struct{
	Id int
	Name string
	Email string
	Password string
	Data T
}

func StructGeneric(){
	u := User[int]{Id: 0,
	Name: "",
Password:"",
Data: 4}
fmt.Println(u)
}
