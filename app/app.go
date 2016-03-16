package app

import (
	"github.com/agent-pink/essocony-go/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"sort"
)

var articles models.Articles
var articleMap = models.ArticleMap{}
var App = mux.NewRouter()

func init() {
	var err error
	articles, err = models.LoadArticles("articles/*.html")
	if err != nil {
		panic(err)
	}
	sort.Sort(articles)
	for _, article := range articles {
		articleMap[article.Slug] = article
	}
	App.HandleFunc("/", Articles)
	App.HandleFunc("/{slug}", Article)
}

var baseTpl = template.Must(template.ParseFiles("templates/base.tpl"))

var articlesTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/style.tpl", "templates/articles.tpl"))

func Articles(w http.ResponseWriter, r *http.Request) {
	articlesTpl.Execute(w, articles)
}

var articleTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/style.tpl", "templates/article.tpl"))

func Article(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	article := articleMap[slug]
	articleTpl.Execute(w, article)
}
