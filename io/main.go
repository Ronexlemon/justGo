package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	r:= strings.NewReader("John Doe The Great Does\n")
	if _,err:= io.Copy(os.Stdout,r);  err !=nil{
		
		log.Fatal(err)
		
	}
	CopyBuffer()
	CopyN()
	


}

//copy buffer

func CopyBuffer()  {
	r1:= strings.NewReader("first String\n")
	r2:= strings.NewReader("Second read\n")

	buf:= make([]byte, 8)
	if _,err:= io.CopyBuffer(os.Stdout,r1,buf);  err !=nil{
		
		log.Fatal(err)
		
	}
	if _,err:= io.CopyBuffer(os.Stdout,r2,buf);  err !=nil{
		
		log.Fatal(err)
		
	}
	
	


	
}

//Copy N

func CopyN(){
	r:=strings.NewReader("Yollow\n")

	result,err:= io.CopyN(os.Stdout,r,0)
	if err!=nil{
		log.Fatal(err)
		}
		fmt.Println("Bytes copied:",result)
}