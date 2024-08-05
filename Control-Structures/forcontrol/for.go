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

func ListNumbers(){
	

	for i:=0;i< num; i++{
		fmt.Printf("Num at Index[%d]:%d\n",i,i)


	}
}

func MuliplicationTable(){

	for i:=1;i<num;i++{
		for y:=1;y<i;y++{
			fmt.Printf("%d * %d = %d\n",i,y,i*y)
		}
		fmt.Println("--------------------------------X",i)
	}
}