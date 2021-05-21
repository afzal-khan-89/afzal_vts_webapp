package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouter() *mux.Router {
	r := mux.NewRouter()

	//fileServer(r) // Fileserver to serve static files
	setWelcomeRouter(r)
	setReportsRouter(r)

	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	return r
}
