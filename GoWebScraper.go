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

  // Create a file to save the scraped data
  file, err := os.Create("scraped_data.txt")
  if err != nil {
    log.Fatal("Error creating file:", err)
  }
  defer file.Close()

  // Find and print all links on the page
  document.Find("a").Each(func(index int, element *goquery.Selection) {
    href, exists := element.Attr("href")
    if exists {
      fmt.Fprintf(file, "Link: %d: %\n", index, href)
    }
  })

  // Find and print all paragraph texts
  document.Find("p").Each(func(index int, element *goquery.Selection) {
    text := strings.TrimSpace(element.Text())
    fmt.Fprintf(file, "Paragraph %d: %s\n", index, text)
  })

  fmt.Println("Scraping completed. Data saved to scraped_data.txt")
}
