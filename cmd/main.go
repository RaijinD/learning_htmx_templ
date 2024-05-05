package main

import (
	"fmt"
	"net/http"

	Components "github.com/raijinD/learning_htmx_templ/cmd/views/components"

	"github.com/a-h/templ"
)

func main() {
	component := Components.Hello("Raijin!")
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.Handle("/", templ.Handler(component))
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Components.Home().Render(r.Context(), w)
			return
		}

	})

	fmt.Printf("Listening on :3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
