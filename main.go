package main

import (
	"github.com/bdcrisman/rest-api-tutorial/articles"
	"github.com/bdcrisman/rest-api-tutorial/handler"
)

func main() {
	articles.Articles = []articles.Article{
		{Title: "Article 1", Desc: "Desc 1", Content: "Content 1"},
		{Title: "Article 2", Desc: "Desc 2", Content: "Content 2"},
	}
	handler.HandleRequests()
}
