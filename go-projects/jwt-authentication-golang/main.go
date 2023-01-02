package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Message struct {
  Status string `json:"status"`
  Info string `json:"info"`
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
  writer.Header().Set("Content-Type", "application/json")
  var message Message
  err := json.NewDecoder(request.Body).Decode(&message)
  if err != nil {
    return
  }
  err = json.NewEncoder(writer).Encode(message)
  if err != nil {
    return
  }
}

func Config(key string) string {
  // load .env file
  err := godotenv.Load(".env")
  if err != nil {
    fmt.Print("Error loading .env file")
  }
  return os.Getenv(key)
}

func generateJWT() (string, error) {
  token := jwt.New(jwt.SigningMethodEdDSA)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(10 * time.Minute)
  claims["authorized"] = true
  claims["user"] = "username"
  
  var sampleSecretKey = []byte(Config("JWT_SECRET"))
  
  tokenString, err := token.SignedString(sampleSecretKey)
  if err != nil {
    return "", err
  }

  return tokenString, nil
}

func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {

  return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

  })

}

func main() {
  /*
  http.HandleFunc("/home", handlePage)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Println("There was an error listening on port :8080", err)
  }
  */
}


