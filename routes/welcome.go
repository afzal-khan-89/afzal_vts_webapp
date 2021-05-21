package routes

import (
	"fmt"
	"net/http"

	"github.com/afzal-khan-89/afzal_vts_webapp/utils"
	"github.com/gorilla/mux"
)

func setWelcomeRouter(r *mux.Router) {
	r.HandleFunc("/", welComeHome).Methods("GET")
	fmt.Print("welcome route set ok   ...")
}

func welComeHome(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "welcome.html", nil)
	fmt.Print("problem here ...")
}
