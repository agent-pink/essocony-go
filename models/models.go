package models

import (
	"bufio"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

type Metadata struct {
	Title  string `json: title`
	Slug   string `json: slug`
	Author string `json: author`
	Date   string `json: date`
	time   *time.Time
}

var houston *time.Location

func init() {
	var err error
	houston, err = time.LoadLocation("America/Chicago")
	if err != nil {
		panic(err)
	}
}

func (m *Metadata) Time() time.Time {
	if m.time != nil {
		return *m.time
	}
	time, err := time.ParseInLocation("2006-01-02 15:04:05", m.Date, houston)
	if err != nil {
		panic(err)
	}
	m.time = &time
	return time
}

type Article struct {
	Metadata
	Contents string
}

func (a *Article) HtmlContents() template.HTML {
	return template.HTML(a.Contents)
}

type ArticleMap map[string]*Article

type Articles []*Article

func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Less(i, j int) bool {
	return a[i].Time().After(a[j].Time())
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func loadArticle(fname string) (Article, error) {
	var article Article
	file, err := os.Open(fname)
	if err != nil {
		return article, err
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	metastr := ""
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		metastr += s.Text() + "\n"
	}
	if s.Err() != nil {
		return article, err
	}
	json.Unmarshal([]byte(metastr), &article.Metadata)
	for s.Scan() {
		article.Contents += s.Text() + "\n"
	}
	if s.Err() != nil {
		return article, err
	}
	return article, err
}
func LoadArticles(glob string) (Articles, error) {
	acc := make(Articles, 0)
	fnames, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	for _, fname := range fnames {
		article, err := loadArticle(fname)
		if err != nil {
			return nil, err
		}
		acc = append(acc, &article)
	}
	return acc, nil
}
