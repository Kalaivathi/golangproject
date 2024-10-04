package main

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
	"golangproject/templates"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ToolBar.html") // Serve the index.html file
	})
	http.HandleFunc("/square", createSquare)
	http.HandleFunc("/rectangle", createRectangle)
	http.HandleFunc("/circle", createCircle)
	http.HandleFunc("/arc", createArc)
	http.HandleFunc("/dropdown", toggleDropdown)
	http.ListenAndServe(":3001", nil)
}

func createSquare(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	component := templates.CreateSquare()
	handler := templ.Handler(component) // To render data on browser
	handler.ServeHTTP(w, r)
}

func createRectangle(w http.ResponseWriter, r *http.Request) {
	component := templates.CreateRectangle()
	handler := templ.Handler(component) // To render data on browser
	handler.ServeHTTP(w, r)
}

func createCircle(w http.ResponseWriter, r *http.Request) {
	component := templates.CreateCircle()
	handler := templ.Handler(component) // To render data on browser
	handler.ServeHTTP(w, r)
}

func createArc(w http.ResponseWriter, r *http.Request) {
	component := templates.CreateArc()
	handler := templ.Handler(component) // To render data on browser
	handler.ServeHTTP(w, r)
}

func toggleDropdown(w http.ResponseWriter, r *http.Request) {
	component := templates.ToggleDropdown()
	handler := templ.Handler(component) // To render data on browser
	handler.ServeHTTP(w, r)
}
