package main

import (
  "fmt"
  "net/http"
  "github.com/auth0/go-jwt-middleware"
  "github.com/dgrijalva/jwt-go"
  "github.com/gorilla/context"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  user := context.Get(r, "user")
  fmt.Fprintf(w, "This is an authenticated request")
  fmt.Fprintf(w, "Claim content:\n")
  for k, v := range user.(*jwt.Token).Claims {
    fmt.Fprintf(w, "%s :\t%#v\n", k, v)
  }
})

func main() {
  jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
      return []byte("secret"), nil
    },
    // When set, the middleware verifies that tokens are signed with the specific signing algorithm
    // If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
    // Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
    SigningMethod: jwt.SigningMethodHS256,
  })

  app := jwtMiddleware.Handler(myHandler)
  http.ListenAndServe("0.0.0.0:3000", app)
}
