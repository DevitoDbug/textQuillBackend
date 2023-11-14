package routes

import (
	"github.com/gorilla/mux"
	"textQuillBackend/pkg/controlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controlers.HandleGenerateText).Methods("POST")
}
