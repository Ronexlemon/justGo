package ifelse

import "fmt"

var num int=100
func IfElse(){
	topic:= "If else ControlStructure"
	fmt.Println(topic)

}

func ListNumbersIfEvenOrOdd(){
	

	for i:=0;i< num; i++{
		if(i%2 ==0){
			fmt.Println(i,"is even")
		}else{
			fmt.Println(i,"is odd")
		}
		// fmt.Printf("Num at Index[%d]:%d\n",i,i)


	}
}