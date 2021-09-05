package postEndpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	a "github.com/bdcrisman/rest-api-tutorial/articles"
)

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var article a.Article
	json.Unmarshal(body, &article)
	a.Articles = append(a.Articles, article)
	json.NewEncoder(w).Encode(article)
}
