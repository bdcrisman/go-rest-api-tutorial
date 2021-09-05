package handler

import (
	"log"
	"net/http"

	get "github.com/bdcrisman/rest-api-tutorial/endpoints/get"
	post "github.com/bdcrisman/rest-api-tutorial/endpoints/post"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// GET
	router.HandleFunc("/", get.HomePage)
	router.HandleFunc("/all", get.ReturnAllArticles)
	router.HandleFunc("/article/{id}", get.ReturnSingleArticle)

	// POST
	router.HandleFunc("/article", post.CreateNewArticle).Methods("POST")

	// PUT

	log.Fatal(http.ListenAndServe(":10000", router))
}
