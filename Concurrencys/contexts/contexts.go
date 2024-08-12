package contexts

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetchAPI(ctx context.Context,url string, results chan <-string){

	req,err:= http.NewRequestWithContext(ctx,"GET",url,nil)

	if err !=nil{
		results <- fmt.Sprintf("Error creating request for %s: %s",url,err.Error())
		return
	}

	client:= http.DefaultClient
	res,err:= client.Do(req)
	if err != nil{
		results <- fmt.Sprintf("Error fetching %s: %s",url,err.Error())
		return
		}
		defer res.Body.Close()

		results <- fmt.Sprintf("Response from %s: %s",url,res.StatusCode)

}


func MainContext(){
	ctx, cancel:= context.WithTimeout(context.Background(),5*time.Second)

	defer cancel()

	urls:=[]string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
		"https://github.com/RonexLemon",
	}

	results:= make(chan string)

	for _,url:= range urls{
		go fetchAPI(ctx,url,results)
	}

	for range urls{
		fmt.Println(<-results)
	}
}