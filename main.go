package main

import (
	"fmt"
	"net/http"

	"github.com/ryanfavour4/gotth-stack-blog/views"
)

func main() {
	fmt.Println("God abeg i came back to this rabbit it does'nt want to leave me alone🐹!")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.Home().Render(r.Context(), w)
	})

	http.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./static")),
	))

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
