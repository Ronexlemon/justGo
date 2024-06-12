package contexts

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type requestid string

var RequestId requestid

func callApi(ctx context.Context, userId int) (bool, error) {
	time.Sleep(400 * time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("context time out exceeded")
	}

	return true, nil
}
func callApi2(ctx context.Context, userId int) (bool, error) {
	time.Sleep(400 * time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("context time out exceeded")
	}

	return true, nil
}

func Contexts() {
	fmt.Println("Starting")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)

	defer cancel()
	userId := 4000
	result, err := callApi(ctx, userId)
	checkNil(err)
	fmt.Println(result)

}

func DoWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Cancelled", ctx.Err())

	}
}

func CallDoWork() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go DoWork(ctx)

	time.Sleep(3 * time.Second)

}

// passwith value
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, RequestId, "12344")
}
func dosomething(ctx context.Context) {
	rID := ctx.Value(RequestId)
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Out of time")
			return
		default:
			fmt.Println("Doing something with", rID)
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func PassWithValue() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)
	go dosomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oooh noo! time exceed")
	}
	time.Sleep(2 * time.Second)
}

// Deadline
func doSomthingWithTimeOut(ctx context.Context) context.Context {

	return context.WithValue(ctx, "yoll", "23456")
}
func returnId(ctx context.Context){
	value := ctx.Value("yoll")
	fmt.Println(value)
}

func WithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = doSomthingWithTimeOut(ctx)

	go returnId(ctx)
	select {
		case <-ctx.Done():
			fmt.Println("Work Done")
		case <- time.After(2*time.Second):
			fmt.Println("time exceed")
			}

			 time.Sleep(3*time.Second)

}

// func WithDeadline(){
// 	ctx,cance:= context.WithDeadline(context.Background(),2*time.)
// }
func WithCancel(){
	ch := make(chan struct{})
	run:= func(ctx context.Context){
		n:=0
		for{
			select{
				case <-ctx.Done():
					fmt.Println("ctx done exiting")
					close(ch)
					return //not to leak the goroutine
				default:
					time.Sleep(time.Millisecond *300)
					fmt.Println(n)
					n++
			}
		}
	}


	ctx,cancel:= context.WithCancel(context.Background())
	go func(){
		time.Sleep(time.Second *2)
		fmt.Println("Good bye")
		cancel()
	}()
	go run(ctx)
	fmt.Println("waiting to cancel...")
	<-ch
	fmt.Println("Bye")

}

//WithTime

func WaitForTZime() {
	ctx,cancel:= context.WithTimeout(context.Background(),1  * time.Second)

	defer cancel()
	fmt.Println(ctx)
}
func checkNil(err error) {
	if err != nil {
		//panic(err)
		log.Fatalf("error is %s", err)
	}
}
