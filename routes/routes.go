package routes

import (
	"log"
	"net/http"

	"github.com/Guilherme129/Primeiro-API-Go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
