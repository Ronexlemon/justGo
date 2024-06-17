package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var indexPage = `<!DOCTYPE html>
<html>
<body>
<h1>My First Web Page</h1>
<p>This is my first web page.</p>
</body>
</html>`

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type userInfo struct {
	username string
	email    string
	age      int
}

type Server struct {
	users map[string]userInfo
}

func New() *Server {
	return &Server{
		users: make(map[string]userInfo),
	}
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexPage))
}

func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodPut:
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("could not read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var u user
		err = json.Unmarshal(body, &u)
		if err != nil {
			log.Printf("could not unmarshal request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("Create user: %v", u.Name)
		s.users[u.Name] = userInfo{
			username: u.Name,
			email:    u.Email,
			age:      u.Age,
		}
		w.WriteHeader(http.StatusCreated)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		u, ok := s.users[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ret := user{
			Name:  name,
			Age:   u.age,
			Email: u.email,
		}

		msg, err := json.Marshal(ret)
		if err != nil {
			log.Printf("could not marshal user: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
