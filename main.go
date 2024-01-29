package main

import (
	"fmt"
	"net/http"
)

const port = ":4269"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Home)
	http.HandleFunc("/easy", Easy)
	http.HandleFunc("/medium", Medium)
	http.HandleFunc("/hard", Hard)
	http.HandleFunc("/rules", Rules)
	fmt.Println("http://localhost:4269 -serve lancé sur" + port)

	http.ListenAndServe(port, nil)
}
