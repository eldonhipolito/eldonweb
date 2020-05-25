package models

import (
	"fmt"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	list := GetAllArticles()

	if len(articleList) != len(list) {
		t.Fail()
	}

	for i, v := range list {

		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}

	}

}

func TestGetArticleByID(t *testing.T) {
	article, err := GetArticleByID(1)

	if err != nil {
		t.Fail()
	}
	fmt.Println(article)

}
