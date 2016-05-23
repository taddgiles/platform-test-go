package main

import (
  "os"
  // "log"

  "github.com/taddgiles/platform-test-go/db"
  "github.com/gin-gonic/gin"
)

const (
  Port = "3000"
)

type User struct {
  Name string
  Email string
}

func init() {
  db.Connect()
}

func main() {
  jwtsecret := os.Getenv("JWT_SECRET")
  if jwtsecret == "" {
    jwtsecret = "secret"
  }

  // db := os.Getenv("MONGODB_URI")
  // if db == "" {
  //   db = "localhost"
  // }

  // session, err := mgo.Dial(db)
  // if err != nil {
  //   panic(err)
  // }
  // defer session.Close()

  // users := session.DB("platform-test").C("users")
  // err = c.Insert(&User{"Ale", "whatnow@email.com"},
  //          &User{"Cla", "whatnow2@email.com"})
  // if err != nil {
  //   log.Fatal(err)
  // }

  // result := User{}

  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ok",
    })
  })

  router.POST("/authenticate", func(c *gin.Context) {
    // email := c.Query("email")
    // password := c.Query("password")

    // err = users.Find(bson.M{"email": String(email)}).One(&result)
    // if err != nil {
    //   log.Fatal(err)
    // }
    // log.Printf("Email:", result)

    c.JSON(200, gin.H{
      "message": "ok",
    })
  })

// Start listening
  port := Port
  if len(os.Getenv("PORT")) > 0 {
    port = os.Getenv("PORT")
  }
  router.Run(":" + port)
}
