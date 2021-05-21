package routes

import (
	"net/http"

	"github.com/afzal-khan-89/afzal_vts_webapp/utils"
	"github.com/gorilla/mux"
)

func setReportsRouter(r *mux.Router) {
	r.HandleFunc("/reports", reportfunc).Methods("GET")

}

func reportfunc(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "reports.html", nil)
}
