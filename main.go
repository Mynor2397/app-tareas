package main

import (
	"log"
	"net/http"

	"github.com/Mynor2397/api-simple/src/routers"
)

func main() {
	router := routers.New()
	if err := http.ListenAndServe(":6060", router); err != nil {
		log.Println("Error al intentar servir: ", err.Error())
	}
}
