package articles

import (
	"testing"
)

func TestPassCreateArticle(t *testing.T) {
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

func TestFailCreateArticle(t *testing.T) {
	a := Article{
		Id:      "",
		Title:   "test title",
		Desc:    "test desc",
		Content: "test content",
	}

	if a.Id == "" || a.Title == "" || a.Desc == "" || a.Content == "" {
		t.Fatalf("Could not create Article")
	}
}
