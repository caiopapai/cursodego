package handlers

import (
	"fmt"
	"net/http"
)

// Index teste
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina de teste")
}
