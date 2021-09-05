package getEndpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	a "github.com/bdcrisman/rest-api-tutorial/articles"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
	fmt.Println("home page endpoint")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllArticles endpoint")
	json.NewEncoder(w).Encode(a.Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range a.Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}
