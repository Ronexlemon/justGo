package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"

	"math/big"
)




func main(){
	welcome:="go cryptography"
	fmt.Println(welcome)
	
	Rand()
	AES()
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
	fmt.Printf("prime numbers %x",n3)
}

func AES(){
	key :=make([]byte,32)
	
	block,err:=aes.NewCipher(key)
	checkNil(err)
	fmt.Printf("%x\n",block)
	
}
func checkNil(e error){
	if e!=nil{
		panic(e)
		}
}