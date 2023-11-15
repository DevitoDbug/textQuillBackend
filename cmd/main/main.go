package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"textQuillBackend/pkg/routes"
)

func main() {
	port := ":8000"

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	fmt.Println("Starting server at port", port)
	http.ListenAndServe(port, r)
}
