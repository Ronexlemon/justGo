package main

import (
	
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	r:= strings.NewReader("John Doe The Great Does\n")
	if _,err:= io.Copy(os.Stdout,r); err !=nil{
		log.Fatal(err)
		
	}
	


}