package binary

import (
	"fmt"
	"io"
	"log"
	"strings"
)




func ReturnBinaryAndAscii(_input string)(string,int){

	r:= strings.NewReader(_input)
	reader:= io.ByteReader(r)

	b,err:= reader.ReadByte()

	if err !=nil{
		log.Fatal(err)
	}
	binaryString:= fmt.Sprintf("%08b", b)
	asciiValue,_:= fmt.Printf("%d",b)
	return binaryString,asciiValue

}