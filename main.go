package main

import (
  "fmt"
  "net/http"
)

func main() {
  request, err := http.NewRequest("GET", "https://blockchain.info/pt/ticker", nil)

  client := http.Client{}
  resp, err := client.Do(request)

  fmt.Println(resp)
  fmt.Println(err)
}
