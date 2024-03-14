package cryptomathrandom

import (
	"fmt"
	"math/rand"
	
)

func CryptoMathRandom() {
	fmt.Println("Welcom to maths in golang")

	//random number
	
	fmt.Println(rand.Intn(5)) //generates a random integer between 0 and 9
}
