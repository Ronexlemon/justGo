package apimock

import (
	"context"
	"fmt"
	"log"
	"time"
)

type response struct{
	userId int
	err error
}

func ApiMock(){
	start := time.Now()
 ctx:= context.Background()
 val,err:=fetchUserData(ctx)
 if err !=nil{
	log.Fatal(err)
 }
 fmt.Printf("User Id: %d\n",val)

	fmt.Println(time.Since(start))


}


//call the api with context
func fetchUserData(ctx context.Context)(int,error){

	ctx,cancel:=context.WithTimeout(ctx,time.Millisecond *250)
	defer cancel()
	responsechanel := make(chan response)
	go func(){
		val,err := apiCallUserData()
		responsechanel <- response{userId:val,err:err}


	}()

	for{
		select {
		case <- ctx.Done():
			return 0,ctx.Err()
		case res:= <- responsechanel:
			return res.userId,res.err
			
		}
	}
}





//mock api call

func apiCallUserData()(int,error){
	time.Sleep(time.Millisecond *200)

	return 8500,nil
}