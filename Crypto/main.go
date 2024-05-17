package main

import (
	"crypto/rand"
	"fmt"
	
	"math/big"
)




func main(){
	welcome:="go cryptography"
	fmt.Println(welcome)
	
	Rand()
}

func Rand(){
	key:=make([]byte,32)
	max := big.NewInt(1000)


	_,err2:= rand.Read(key)
	checkNil(err2)
	n2,err:= rand.Int(rand.Reader,max)
	checkNil(err)

	n3,err:= rand.Prime(rand.Reader,7)
	checkNil(err)
	


	fmt.Printf("Key: %x\n", key)
	fmt.Println(n2)
	fmt.Println("prime numbers",n3)
}
func checkNil(e error){
	if e!=nil{
		panic(e)
		}
}