package server

import "net/http"

var indexPage =`<!DOCTYPE html>
<html>
<body>
<h1>My First Web Page</h1>
<p>This is my first web page.</p>
</body>
</html>`

var userInfo =`
{
"name":"John doe",
"age":45}`

type Server struct{

}

func New() *Server{
	return &Server{}
}


func(s *Server) HandleIndex(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexPage))

}

func(s *Server) HandleUser(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userInfo))

}