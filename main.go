package main

import (
	"fmt"
	"net/http"
)

const port = ":4269"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/game", Game)
	fmt.Println("http://localhost:4269 -serve lancé sur" + port)

	http.ListenAndServe(port, nil)
}
