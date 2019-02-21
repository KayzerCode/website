package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", hom)
	http.HandleFunc("/about", abou)
	http.HandleFunc("/nodes", nods)
	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func hom(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "default.gohtml", nil)

	//TODO check if template file exists
}

func nods(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "nodes.gohtml", nil)
}

func abou(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
