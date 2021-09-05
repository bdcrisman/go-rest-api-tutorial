package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bdcrisman/rest-api-tutorial/articles"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
	fmt.Println("home page endpoint")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllArticles endpoint")
	json.NewEncoder(w).Encode(articles.Articles)
}
