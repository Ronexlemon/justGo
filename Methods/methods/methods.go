package methods

import "fmt"


type Rectangle struct{
	width, height float64
}


func (r *Rectangle) Square()float64{
	return r.width * r.height
}
func (r *Rectangle) Double(f float64){
	r.width *= f
	r.height *= f
}

func Rec(){
	r:= Rectangle{10,5}
	r.Double(10)
	
	fmt.Println(r.Square())
}