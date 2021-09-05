package articles

import (
	"testing"
)

func TestCreateArticle(t *testing.T) {
	a := Article{
		Id:      "1",
		Title:   "test title",
		Desc:    "test desc",
		Content: "test content",
	}

	if a.Id == "" || a.Title == "" || a.Desc == "" || a.Content == "" {
		t.Fatalf("Could not create Article")
	}
}
