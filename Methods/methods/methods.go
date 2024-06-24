package methods

import "fmt"


type Rectangle struct{
	width, height float64
}


func (r *Rectangle) Square()float64{
	return r.width * r.height
}

func Rec(){
	r:= Rectangle{10,5}
	fmt.Println(r.Square())
}