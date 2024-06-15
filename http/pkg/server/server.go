package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var indexPage =`<!DOCTYPE html>
<html>
<body>
<h1>My First Web Page</h1>
<p>This is my first web page.</p>
</body>
</html>`
//user represents the JSON value thats send a response to a request
type user struct{
	Name string `json:"name`
	Email string `json:"email"`
	Age int `json:"age"`
}

//information that is stored per user
type userInfo struct{
	username string
	age int
}

type Server struct{
	users map[string]userInfo //key -> username

}

func New() *Server{
	return &Server{
		users: make(map[string]userInfo),
	}
}


func(s *Server) HandleIndex(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexPage))

}
func(s *Server) HandleIndexe(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexPage))

}

func(s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodPost , http.MethodPut:
		if contentType := r.Header.Get("Content-Type");contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
	
		 body,err := io.ReadAll(r.Body)
		 
		 if err != nil{
			log.Printf("could not read request body : %v",err)
			w.WriteHeader(http.StatusInternalServerError)
			return
			}
			defer r.Body.Close()
			//unMarshall
			var u user
			errr:= json.Unmarshal(body,&u)
			if errr !=nil{
				log.Printf("could not unmarshal request body : %v",err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Printf("Create user : %v",u.Name)
			s.users[u.Name] = userInfo{
				username: u.Name,
				age: u.Age,
			}
			

		

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
	

}