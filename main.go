package main

import (
	"github.com/bdcrisman/rest-api-tutorial/articles"
	handler "github.com/bdcrisman/rest-api-tutorial/handlers"
)

func main() {
	articles.Articles = []articles.Article{
		{Id: "1", Title: "Article 1", Desc: "Desc 1", Content: "Content 1"},
		{Id: "2", Title: "Article 2", Desc: "Desc 2", Content: "Content 2"},
	}
	handler.HandleRequests()
}
