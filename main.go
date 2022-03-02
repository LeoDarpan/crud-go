package main

import (
	"CRUD/models"
	"mux/routes"
	"mux/utils"

	"net/http"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")

	//Initiating Router
	r := routes.NewRouter()

	//Telling the server to use router for all the routes starting with '/'
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
