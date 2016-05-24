package main

import (
  "os"
  "log"

  "database/sql"
  _ "github.com/lib/pq"
  "github.com/gin-gonic/gin"
)

const (
  Port = "3000"
)

func main() {
  dbUrl := os.Getenv("DATABASE_URL")
  if dbUrl == "" {
    dbUrl = "postgres://localhost/platform_test_development"
  }

  db, err := sql.Open("postgres", dbUrl + "?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  jwtsecret := os.Getenv("JWT_SECRET")
  if jwtsecret == "" {
    jwtsecret = "secret"
  }


  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ok",
    })
  })

  router.GET("/api/v1/users/current", func(c *gin.Context) {
    id := 1

    var email string
    var name string
    err := db.QueryRow("SELECT email,name FROM users WHERE id=$1", id).Scan(&email, &name)
    switch {
    case err == sql.ErrNoRows:
      c.JSON(401, gin.H{
        "message": "No user with that ID",
      })
    case err != nil:
      log.Fatal(err)
    default:
      c.JSON(200, gin.H{
        "email": email,
        "name": name,
      })
    }
  })

// Start listening
  port := Port
  if len(os.Getenv("PORT")) > 0 {
    port = os.Getenv("PORT")
  }
  router.Run(":" + port)
}
