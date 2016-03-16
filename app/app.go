package app

import (
	"github.com/agent-pink/essocony-go/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"sort"
)

var articles models.Articles
var App = mux.NewRouter()

func init() {
	var err error
	articles, err = models.LoadArticles("articles/*.html")
	if err != nil {
		panic(err)
	}
	sort.Sort(articles)
	App.HandleFunc("/", Articles)
}

var baseTpl = template.Must(template.ParseFiles("templates/base.tpl"))

var articlesTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/articles.tpl"))

func Articles(w http.ResponseWriter, r *http.Request) {
	articlesTpl.Execute(w, articles)
}
