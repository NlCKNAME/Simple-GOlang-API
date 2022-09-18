package main

//import of used packages
import (
	"net/http"

	"./myserver"
)

func main() {
	server, err := myserver.NewServer()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", server)
}
