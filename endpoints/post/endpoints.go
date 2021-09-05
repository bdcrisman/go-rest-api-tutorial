package postEndpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(body))
}
