package main

import (
	"fmt"
	"net/http"

	"github.com/afzal-khan-89/afzal_vts_webapp/routes"
	"github.com/afzal-khan-89/afzal_vts_webapp/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("server starting ... ... ...")

	utils.LoadTemplate()
	routes.SetRouter()

	http.ListenAndServe(":8080", nil)
}
