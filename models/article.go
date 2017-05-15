package models

import "errors"

type Article struct {
  Id      int     `json:"id"`
  Title   string  `json:"title"`
  Content string  `json:"content"`
}

var articleList = []Article{
  Article{Id: 1, Title: "Article 1", Content: "Article 1 body"},
  Article{Id: 2, Title: "Article 2", Content: "Article 2 body"},
  Article{Id: 3, Title: "Article 3", Content: "Article 3 body"},
}

func GetAllArticles() []Article {
  return articleList
}

func GetArticleById(id int) (*Article, error) {
  for _, a := range articleList {
    if a.Id == id {
      return &a, nil
    }
  }
  return nil, errors.New("Article not found")
}
