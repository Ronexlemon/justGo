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

	checkNilError(err)
	length,err:=io.WriteString(file,content)
	checkNilError(err)
	fmt.Println("len is:",length)

	defer file.Close()

	ReadFiles("./files/myfile.txt")

}

//reading the file

func  ReadFiles(filename string){
	databyte,err:=os.ReadFile(filename) //read into the bytes format
	checkNilError(err)                   //checking if there was
		
		data:=string(databyte)
		
		fmt.Println(data)


}

func checkNilError(err error){
	if err!= nil {	
		panic(err)}

}