package slices

import ("fmt"
"slices")

var Slice=[]int{1,2,3}
	var Slice2 = append(Slice, 4,5,6)
	var Slice3 = []int{1,2,3}
//dynamic slice
var integers []int
func InitSlice(){
	//using append
	integers = append(integers, 100)
	integers = append(integers, 200)
	integers = append(integers, 300)
	//integers[40] = 4000
	fmt.Println("slice",integers)


}

func InitSlice2(){
	//using make
	integers := make([]int, 5,8)
	integers[0] = 100
	integers[1] = 400
	integers[2] = 800
	integers[3] = 900
	integers[4] = 1000
	fmt.Println("slice",integers)


}

func AppendTo(){
	slice:=[]int{1,2,3}
	slice = append(slice, 4,5,6)
	fmt.Println("new Appended Slice",slice)
}

func Copy(){
	slice1:=[]int{1,2,3}
	slice2:=make([]int, len(slice1))
	 copy(slice2,slice2)
	 fmt.Println("copied slice",slice2)
	 fmt.Println("updated slice",slice1)

}

func Compareslices(slice1 []int,slice2[]int)int{
	return slices.Compare(slice1,slice2)
	
	
}

