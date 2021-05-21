package utils

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

//LoadTemplate ...
func LoadTemplate() {
	//tpl = template.Must(template.ParseGlob("/home/babu/work-zone/geon/geo-web/templates/*"))  working :)
	//tpl = template.Must(template.ParseGlob("./templates/*")) //working and looking good :)
	//tpl = template.Must(template.ParseGlob("templates/*")) //working :)

	tpl = template.Must(template.ParseFiles(
		"./templates/welcome/index.html",
		"./templates/reports/reports.html",
		"./templates/header/header.html",
		"./templates/navbar/navbar.html",
	))
}

//ExecuteTemplate ...
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := tpl.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		log.Print("error to execute template " + tmpl)
	}
}
