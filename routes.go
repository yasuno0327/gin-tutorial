package main

import (
  handler "./handlers"
  "github.com/gin-gonic/gin"
)

func initializeRoutes() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  router.GET("/", handler.ShowIndexPage)
  router.GET("/article/view/:article_id", handler.GetArticle)
  router.Run(":8080")
}
