package handler

import (
	"log"
	"net/http"

	"github.com/bdcrisman/rest-api-tutorial/endpoints"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", endpoints.HomePage)
	router.HandleFunc("/all", endpoints.ReturnAllArticles)

	log.Fatal(http.ListenAndServe(":10000", router))
}
