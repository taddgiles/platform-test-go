package main

import (
  "os"
  "github.com/gin-gonic/gin"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }


  r := gin.New()
  r.Use(gin.Logger())

  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ok",
    })
  })

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  r.Run(":" + port)
}
