package forcontrol

import "fmt"

var num int = 100;

func ListNumber(){
	i:=0
	
for  i< num{
	fmt.Printf("Num at Index[%d]:%d\n",i,i)

	i+=1
}

	
}