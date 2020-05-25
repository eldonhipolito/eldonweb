package models

import (
	"errors"
	"sort"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []Article {
	return articleList
}

func SetAllArticles(list []Article) {
	articleList = list
}

func GetArticleByID(id int) (Article, error) {
	index := sort.Search(len(articleList), func(i int) bool { return articleList[i].ID >= id })

	if index == len(articleList) {
		return Article{}, errors.New("Article not found")
	}

	return articleList[index], nil
}
