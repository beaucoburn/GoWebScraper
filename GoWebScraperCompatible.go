package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "strings"
)

func main() {
  // URL to scrape
  url := "https://dev.to/beaucoburn/should-i-choose-tailwind-or-vanilla-css-ccl"

  // Make HTTP GET request
  response, err := http.Get(url)
  if err != nil {
    log.Fatal("Error fetching the URL:", err)
  }
  defer response.Body.Close()

  if response.StatusCode != 200 {
    log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status)
  }

  // Read the response body
  body, err := io.ReadAll(response.Body)
  if err != nil {
    log.Fatal("Error reading HTTP response body:", err)
  }

  // Convert body to string
  content := string(body)

  // Create a file to save the scraped data
  file, err := os.Create("scraped_data.txt")
  if err != nil {
    log.Fatal("Error creating file:", err)
  }
  defer file.Close()

  // Simple (but imperfect) way to find links
  links := strings.Split(content, "<a href=\"")
  for i, link := range links {
    if i == 0 { // Skip the part before the first link
      continue
    }
    end := strings.Index(link, "\"")
    if end != -1 {
      fmt.Fprintf(file, "Link %d: %s\n", i, link[:end])
    }
  }

  // Simple (but imperfect) way to find paragraph texts
  paragraphs := strings.Split(content, "<p>")
  for i, p := range paragraphs {
    if i == 0 { // Skip the part before the first paragraph
      continue
    }
    end := strings.Index(p, "</p>")
    if end != -1 {
      text := strings.TrimSpace(p[:end])
      fmt.Fprintf(file, "Paragraph %d: %s\n", i, text)
    }
  }

  fmt.Println("Scraping complete. Data saved to scraped_data.txt")
}
