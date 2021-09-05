package handler

import (
	"log"
	"net/http"

	delete "github.com/bdcrisman/rest-api-tutorial/endpoints/delete"
	get "github.com/bdcrisman/rest-api-tutorial/endpoints/get"
	post "github.com/bdcrisman/rest-api-tutorial/endpoints/post"
	put "github.com/bdcrisman/rest-api-tutorial/endpoints/put"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// GET
	router.HandleFunc("/", get.HomePage)
	router.HandleFunc("/all", get.ReturnAllArticles).Methods("GET")
	router.HandleFunc("/article/{id}", get.ReturnSingleArticle).Methods("GET")

	// POST
	router.HandleFunc("/article", post.CreateNewArticle).Methods("POST")

	// DELETE
	router.HandleFunc("/article/{id}", delete.DeleteArticle).Methods("DELETE")

	// PUT
	router.HandleFunc("/article/{id}", put.UpdateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", router))
}
