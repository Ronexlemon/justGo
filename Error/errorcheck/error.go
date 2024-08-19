package errorcheck

import (
	"fmt"
	"log"
	"os"
)

func OpenAfile(flename string) {
	file, err := os.ReadFile(flename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("file content : %s\n",file)
	

}


