package conversion

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Conversion() {
	welcome := "Welcome to lemonr hacks"
	rateMessage := "Please rate our hacks fbtwn 1  & 5"
	fmt.Println(welcome)
	fmt.Println(rateMessage)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thank you for rating us a", input)
	//conversion
	numRating,err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil{
		//panic(err)
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to rating",numRating+1)
	}

}
