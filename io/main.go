package main

import (
	"fmt"
	// "goio/binary"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("John Doe The Great Does\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {

		log.Fatal(err)

	}
	// CopyBuffer()
	// CopyN()
	// Pipe()
	// ReadAll()
	// ReadAtleast()
	// ReadFull()
	// WriteString()
	//Writer()
	//Reader()
	//ByteReader()
	//ByteScanner()
	// binaryForm,intform:=binary.ReturnBinaryAndAscii("H")
	// fmt.Println(binaryForm)
	// fmt.Println(intform)
	Seeker()

}

//copy buffer

func CopyBuffer() {
	r1 := strings.NewReader("first String\n")
	r2 := strings.NewReader("Second read\n")

	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {

		log.Fatal(err)

	}
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {

		log.Fatal(err)

	}

}

//Copy N

func CopyN() {
	r := strings.NewReader("Yollow\n")

	result, err := io.CopyN(os.Stdout, r, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bytes copied:", result)
}

//Go pipe

func Pipe() {
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		fmt.Fprint(w, "Hello Pipe\n")

	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

// io ReadAll
func ReadAll() {
	r := strings.NewReader("Hello ReadAll screw it\n")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

// Read Atleast

func ReadAtleast() {
	/*
	*func ReadAtLeast(r Reader,[]byte,min int)(n int,error)
	*The error is EOF only if no bytes were read.
	*if an EOF happens after reading  fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF
	*if min is greater than the length of buf,ReadAtLeast returns ErrShortBuffer.
	*on Return , n>= min if and only if err ==nil
	*if r returns an error having read at least min bytes, the error is droped
	 */
	r := strings.NewReader("Hello ReadAtleast screw it\n")
	buf := make([]byte, 1)

	resultint, err := io.ReadAtLeast(r, buf, 1)

	if err == io.ErrShortBuffer {
		log.Fatal(err)
	}

	fmt.Println(resultint)

}

// ReadFull
func ReadFull() {
	r := strings.NewReader("Hello ReadFull screw it\n")
	buf := make([]byte, 10)
	resultint, err := io.ReadFull(r, buf)
	if err != nil {
		log.Fatal(err)
	}
	Print(resultint)

}

// Writestring
func WriteString() {
	if _, err := io.WriteString(os.Stdout, "Yollow writeString"); err != nil {
		log.Fatal(err)
	}
}

//interface
//Write

func Writer() {
	file, _ := os.Create("./chat.txt")
	writer := io.Writer(file)
	writer.Write([]byte("Hello Writer"))

	defer file.Close()

}

func Reader() {
	
	file, _ := os.Open("./chat.txt")
	reader := io.Reader(file)
	buf := make([]byte, 1)
	for{
		n, err := reader.Read(buf)
		fmt.Println(n)

		if err !=nil{
			break
		}
	}
	defer file.Close()
	
	
}
//Seeker 

func Seeker(){
	file,_:=os.Open("./chat.txt")
	reader:= io.Reader(file)
	
	buffer,err:= io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v} \n",string(buffer),err)
	seeker:= reader.(io.Seeker)
	seeker.Seek(0,io.SeekStart)
	buffer,err= io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v} \n",string(buffer),err)

	

	if err!=nil{
		log.Fatal(err)
	}
	

}

// Types
// ByteReader
func ByteReader() {
	r := strings.NewReader("7")

	reader := io.ByteReader(r)
	b, err := reader.ReadByte()
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("in binary form %b  \n", b)
	fmt.Printf("in integer form %d  \n", b)
}

// ByteScanner
func ByteScanner() {
	r := strings.NewReader("Hello BytScanner")
	reader := io.ByteScanner(r)
	err := reader.UnreadByte().Error()
	fmt.Println(err)

	for {
		b, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			fmt.Println("error", err)
			break
		}
		fmt.Printf("in binary form %b  \n", b)
		fmt.Printf("in integer form %d  \n", b)
	}

}

func checkNill(err error) {
	log.Fatal(err)
}
func Print(input interface{}) {
	fmt.Println(input)
}