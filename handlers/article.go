package handlers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
  models "../models"
)

func ShowIndexPage(c *gin.Context) {
  articles := models.GetAllArticles()

  c.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "title": "Home Page",
      "payload": articles,
    },
  )
}

func GetArticle(c *gin.Context) {
  if articleId, err := strconv.Atoi(c.Param("article_id")); err == nil {
    if article, err := models.GetArticleById(articleId) ; err == nil {
      c.HTML(
        http.StatusOK,
        "article.html",
        gin.H{
          "title": article.Title,
          "payload": article,
        },
      )
    } else {
      c.AbortWithError(http.StatusNotFound, err)
    }
  } else {
    c.AbortWithStatus(http.StatusNotFound)
  }
}
