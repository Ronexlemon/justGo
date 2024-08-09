package interfaces

import (
	"fmt"
	"log"
	"strconv"
)
type Book struct{

	Author string
	Title string
}

func (b Book) String()string{

	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int
func (c Count)String()string{

	return strconv.Itoa(int(c))

}


func WriteLog(s fmt.Stringer){
	log.Print(s.String())
}

func GetLogs(){
	book:= Book{Title: "The art of acting",Author: "John Doe"}
	WriteLog(book)
	count:= Count(3)
	WriteLog(count)
}

type CheckReceipt struct{
	Amount float64
	Receipt string
}

type I interface{
	Check()

}


func (c *CheckReceipt) Check(){
	if c.Amount > 0{
		fmt.Println("Receipt is valid and number is:",c.Receipt)
		
	}else{
		fmt.Println("Receipt is invalid")
	}
	


}

func Receipt(){
	var rec I = &CheckReceipt{Amount: 10,Receipt: "1000856T"}
	rec.Check()
}