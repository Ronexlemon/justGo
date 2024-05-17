package contexts

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)



func callApi(ctx context.Context, userId int)(bool,error){
	time.Sleep(400 *time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded{
		return false, errors.New("context time out exceeded")
	}

	return true, nil
}

func Contexts(){
	fmt.Println("Starting")
ctx,cancel := context.WithTimeout(context.Background(),500*time.Millisecond)

defer cancel()
userId:=4000
result,err := callApi(ctx,userId)
checkNil(err)
fmt.Println(result)

	
}

func checkNil(err error){
	if err !=nil{
		//panic(err)
		log.Fatalf("error is %s",err)
	}
}
