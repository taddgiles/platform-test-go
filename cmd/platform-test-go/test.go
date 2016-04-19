package main

import (
  "fmt"
  "log"
  "os"
  // "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  Name string
  Email string
}

func main() {

  jwtsecret := os.Getenv("JWT_SECRET")
  if jwtsecret == "" {
    jwtsecret = "secret"
  }
  fmt.Println(jwtsecret)

  // r := gin.Default()
  // r.GET("/api/v1/users", func(c *gin.Context) {
  //   c.JSON(200, gin.H{
  //     "message": "pong",
  //   })
  // })
  // r.Run() // listen and server on 0.0.0.0:8080

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("platform-test").C("users")
  err = c.Insert(&User{"Ale", "whatnow@email.com"},
           &User{"Cla", "whatnow2@email.com"})
  if err != nil {
    log.Fatal(err)
  }

  result := User{}
  err = c.Find(bson.M{"email": "whatnow@email.com"}).One(&result)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Email:", result)
}
