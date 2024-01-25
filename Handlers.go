package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
	RenderTemplate(w, "Home")
}
func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, game!")
	RenderTemplate(w, "Game")
}
func RenderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./Tempates/" + html + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}
