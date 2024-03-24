package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var ola [6]string
	ola[0] = "Oi "
	ola[1] = "bom "
	ola[2] = "dia "
	ola[3] = "tudo "
	ola[4] = "bem "
	ola[5] = "? "
	fmt.Fprint(w, "<h1>"+ola[0]+ola[1]+ola[2]+ola[3]+ola[4]+ola[5]+"</h1")
}
