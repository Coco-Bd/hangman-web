package main

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "acceuil")
}
func Easy(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "easy")
}
func Medium(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "medium")
}
func Hard(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "hard")
}
func Rules(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "rules")
}
func RenderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./html/" + html + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}
