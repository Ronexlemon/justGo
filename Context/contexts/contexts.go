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
func callApi2(ctx context.Context, userId int)(bool,error){
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


func DoWork(ctx context.Context){
	select{
	case <-time.After(2* time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Cancelled",ctx.Err())

	}
}

func CallDoWork(){
	ctx,cancel:=context.WithTimeout(context.Background(),1* time.Second)
	defer cancel()

	go DoWork(ctx)

	time.Sleep(3 *time.Second)

}

func checkNil(err error){
	if err !=nil{
		//panic(err)
		log.Fatalf("error is %s",err)
	}
}
