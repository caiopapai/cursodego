package main

import (
	"fmt"
	"net/http"

	"github.com/caiopapai/curso_de_go/server/handlers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ola Mundo")
	})

	http.HandleFunc("/teste", handlers.Index)

	http.ListenAndServe(":1313", nil)
}
