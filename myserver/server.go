package myserver

//import of used packages
import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var Version = "v0.1.5"

//Data structure to stock MySQL's results
type User struct {
	ID  int    `json:"ID"`
	Nom string `json:"Nom"`
}

var user User

var users []User

// Server is the struct who hold all we need to run the service
type Server struct {
	// Http router
	chi.Router

	// UserList
	users []User
}

// NewServer is the default call to create a proxy (Abstraction)
func NewServer() (server *Server, err error) {
	return NewServerParams()
}

//Main Function
func NewServerParams() (rtr *Server, err error) {

	// Create chi router
	router := chi.NewRouter()

	// Integrate it in our server service
	rtr = &Server{Router: router}
	//Creation of new interface
	rtr.Use(middleware.RequestID)
	rtr.Use(middleware.Logger)
	rtr.Use(middleware.Recoverer)
	rtr.Use(middleware.URLFormat)
	rtr.Use(render.SetContentType(render.ContentTypeJSON))

	//GET request if url == "X.X.X.X:3334/version"
	rtr.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"Version\":\"" + Version + "\"}"))
	})

	//GET request if url == "X.X.X.X:3334/users"

	rtr.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(rtr.users)

		if err != nil {
			w.WriteHeader(500)
			return
		}
	})

	rtr.Get("/users/{UserID}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		for i, a := range rtr.users {
			if strconv.Itoa(a.ID) == chi.URLParam(r, "UserID") {
				err := json.NewEncoder(w).Encode(rtr.users[i])
				if err != nil {
							// http 500 so clean... such wow !
					w.WriteHeader(500)
							// Would be better with an error message but well... we'll see later :)
					return
				}
			}
		}
	})

	rtr.Post("/NewUser/user={UserName}/id={UserID}/", func(w http.ResponseWriter, r *http.Request) {
		UserName := chi.URLParam(r, "UserName")
		UserID := chi.URLParam(r, "UserID")
		if UserName == "" {
			w.WriteHeader(400)
			return
		}
		
		if UserID == "" {
			w.WriteHeader(400)
			return
		}

		ID, err := strconv.Atoi(UserID)
		if err != nil {
			panic(err)
		}

		var Chk bool = rtr.ChkUser(ID)

		if Chk == false {
			var user = User{ID: ID, Nom: UserName}
			rtr.AddUser(user)
			w.Header().Set("Content-Type", "application/json")
			for i, a := range rtr.users {
				err := json.NewEncoder(w).Encode(rtr.users[i])
				if err != nil {
							// http 500 so clean... such wow !
					w.WriteHeader(500)
							// Would be better with an error message but well... we'll see later :)
					return
				}
				fmt.Println(a)
			}
		}else{
			w.WriteHeader(403)
			return
		}
	})

	//GET request if we whant to create a User

	rtr.Delete("/delete/{UserID}", func(w http.ResponseWriter, r *http.Request) {
		if UserID := chi.URLParam(r, "UserID"); UserID != "" {

			var Pos int = 0
			var Chk bool = false

			//Get the position of the user with the ID
			for i, a := range rtr.users {
				if strconv.Itoa(rtr.users[i].ID) == UserID {
					Pos = i
					Chk = true
					if rtr.users == nil {
						panic(a)
					}
				}
			}

			if Chk == true {
				rtr.DeleteUser(Pos)
				w.Write([]byte("User has been deleted !"))
			} else {
				w.Write([]byte("User not exist !"))
			}
		}
	})

	//Port on wich the server listen
	return
}

// AddUser adds a new user in our Server
func (s *Server) AddUser(newuser User) (err error) {
	// TODO : Code Code
	s.users = append(s.users, newuser)
	s.sort_users()
	return
}

// DeleteUser delete a user in our Server
func (s *Server) DeleteUser(userid int) (err error) {
	// TODO : Code Code

	var Pos int = 0
	var Chk bool = false

	//Get the position of the user with the ID
	for i, a := range s.users {
		if s.users[i].ID == userid {
			Pos = i;
			Chk = true
			if s.users == nil {
				panic(a)
			}
		}
	}

	if Chk == true {
		s.users[len(s.users)-1], s.users[Pos] = s.users[Pos], s.users[len(s.users)-1]
		s.users = s.users[:len(s.users)-1]
		s.sort_users()
	}
	if Chk == false {

	}

	return
}

// sort_users users of our server
func (s *Server) sort_users() (err error) {
	// TODO : Code Code

	var User_Mem User
	var Ctl int = 0

	for Ctl < len(s.users) {
		var Mem int = Ctl
		var h int = Ctl
		for h, a := range s.users {
			if s.users[h].ID > s.users[Mem].ID {
				User_Mem = s.users[Mem]
				s.users[Mem] = s.users[h]
				s.users[h] = User_Mem
				if s.users == nil {
					panic(a)
				}
			}
		}
		if s.users == nil {
			panic(h)
		}
		Ctl = Ctl + 1
	}
	return
}

func (s *Server) ChkUser(userid int) bool {
	var Chk bool = false

	for i, a := range s.users {
		if s.users == nil {
			panic(i)
		}
		if a.ID == userid {
			Chk = true
		} else {
			Chk = false
		}
	}
	return Chk
}