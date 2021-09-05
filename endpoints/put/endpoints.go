package postEndpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	a "github.com/bdcrisman/rest-api-tutorial/articles"
	"github.com/gorilla/mux"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating article...")

	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range a.Articles {
		if article.Id == id {
			var updatedArticle a.Article
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &updatedArticle)
			article = updatedArticle
			json.NewEncoder(w).Encode(article)
		}
	}
}
