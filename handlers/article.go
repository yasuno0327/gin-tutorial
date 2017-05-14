package handlers

import (
  "net/http"
  "github.com/gin-gonic/gin"
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
