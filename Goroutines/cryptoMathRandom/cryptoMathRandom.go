package cryptomathrandom

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func CryptoMathRandom() {
	fmt.Println("Welcom to maths in golang")

	//random number
	
	//fmt.Println(rand.Intn(5)) //generates a random integer between 0 and 9
}


func CryptoRandom(){
	myRandomNum,_:= rand.Int(rand.Reader,big.NewInt(5))

	fmt.Println(myRandomNum)

}