package files

import (
	"fmt"
	"io"
	
	"os"
)

func Files(){
	welcome:="Welcome to files in golang with @lemonr"
	fmt.Println(welcome)

	content:="This need to go in a file as simple as that -> lemonr.io"

	file,err:=os.Create("./files/myfile.txt");

	if err !=nil{
		panic(err);
	}
	length,err:=io.WriteString(file,content)
	if err !=nil{
		panic(err);
	}
	fmt.Println("len is:",length)

	defer file.Close()

	ReadFiles("./files/myfile.txt")

}

//reading the file

func  ReadFiles(filename string){
	databyte,err:=os.ReadFile(filename) //read into the bytes format
	if err!= nil {	
		panic(err)}
		
		data:=string(databyte)
		
		fmt.Println(data)


}