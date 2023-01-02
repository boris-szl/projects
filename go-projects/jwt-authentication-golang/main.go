package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
  err := json.NewEncoder(writer).Encode(message)
  if err != nil {
    return
  }
}

func main() {
  http.HandleFunc("/home", handlePage)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Println("There was an error listening on port :8080", err)
  }

}




