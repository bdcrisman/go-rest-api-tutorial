package articles

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func NewArticle(id, title, desc, content string) Article {
	return Article{
		Id: id, Title: title, Desc: desc, Content: content,
	}
}

var Articles []Article
