package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "strings"

  "github.com/PuerkitoBio/goquery"
)

func main() {
  fmt.Println("Go Web Scraper")

  url := "http://example.com"

  response, err := http.Get(url)
  if err != nil {
    log.Fatal("Error fetching the URL:", err)
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status )
  }

  // Create a goquery document from the HTTP response
  document, err := goquery.NewDocumentFromReader(response.Body)
  if err != nil {
    log.Fatal("Error loading HTTP response body:", err)
  }


}
