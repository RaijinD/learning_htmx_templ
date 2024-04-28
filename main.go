package main

import (
	"fmt"
	Components "htmx_templ/view"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := Components.Hello("Raijin")

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
