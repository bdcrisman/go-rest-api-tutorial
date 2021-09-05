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
			body, _ := ioutil.ReadAll(r.Body)
			var updatedArticle a.Article
			json.Unmarshal(body, &updatedArticle)
			article = updatedArticle
			json.NewEncoder(w).Encode((article))
			break
		}
	}

	// body, _ := ioutil.ReadAll(r.Body)
	// var article a.Article
	// json.Unmarshal(body, &article)
	// a.Articles = append(a.Articles, article)
	// json.NewEncoder(w).Encode(article)
}
