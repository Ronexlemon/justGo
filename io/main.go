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
	Pipe()
	ReadAll()
	ReadAtleast()
	


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

//Go pipe

func Pipe(){
	r,w:=io.Pipe()

	go func ()  {
		defer w.Close()
		fmt.Fprint(w,"Hello Pipe\n")
		
	}()

	if _,err:=io.Copy(os.Stdout,r); err !=nil{
		log.Fatal(err)
	}
}

//io ReadAll
func ReadAll(){
	r:=strings.NewReader("Hello ReadAll screw it\n")

	b,err:= io.ReadAll(r)
	if err!=nil{
			log.Fatal(err)}
	fmt.Println(b)
}

// Read Atleast

func ReadAtleast(){
	/*
	*func ReadAtLeast(r Reader,[]byte,min int)(n int,error)
	*The error is EOF only if no bytes were read.
	*if an EOF happens after reading  fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF
	*if min is greater than the length of buf,ReadAtLeast returns ErrShortBuffer.
	*on Return , n>= min if and only if err ==nil
	*if r returns an error having read at least min bytes, the error is droped
	*/
	r:=strings.NewReader("Hello ReadAtleast screw it\n")
	buf:=make([]byte,1)

	resultint,err:=io.ReadAtLeast(r,buf,1)
	if err == io.ErrShortBuffer{
		log.Fatal(err)
	}

	fmt.Println(resultint)

}

func checkNill(err error){
	log.Fatal(err)
}

