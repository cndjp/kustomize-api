package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// route to kustomize
	r.Path("/kustomize").
		Methods("POST").
		HandlerFunc(KustomizeHandler)

	http.ListenAndServe(":8080", r)
}

func KustomizeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)

	fmt.Fprintf(w, "body: %v", buffer.String())
}
