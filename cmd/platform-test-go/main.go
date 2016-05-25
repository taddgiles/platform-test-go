package main

import (
  "os"
  "log"
  "strings"

  "database/sql"
  _ "github.com/lib/pq"
  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"
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

  router.GET("/loaderio-c6b9e7b90d19f9d6123be11389320bcc/", func(c *gin.Context) {
    c.String(200, "loaderio-c6b9e7b90d19f9d6123be11389320bcc")
  })

  router.GET("/api/v1/users/current", func(c *gin.Context) {
    authorizationHeader := c.Request.Header.Get("Authorization")
    bearerToken := strings.Split(authorizationHeader, " ")[1]

    var user_id string

    token, jwterr := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
      return []byte(jwtsecret), nil
    })

    if jwterr == nil && token.Valid {
      user_id = token.Claims["user_id"].(string)
    } else {
      c.JSON(401, gin.H{
        "message": "Invalid",
      })
    }

    var email string
    var name string
    var created_at string
    var updated_at string
    err := db.QueryRow("SELECT email,name,created_at,updated_at FROM users WHERE id=$1", user_id).Scan(&email, &name, &created_at, &updated_at)
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
        "created_at": created_at,
        "updated_at": updated_at,
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
