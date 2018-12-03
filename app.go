package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()

	router.GET("/ping", Ping)

	log.Fatal(http.ListenAndServe(":80", (router)))
}

func Ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("p")
	return
}
