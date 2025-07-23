# 🕷️ GoScraper – A Powerful Go Web Scraping Package

GoScraper is a lightweight but advanced Go module for web scraping. It helps developers extract key information from web pages such as metadata, links, images, and also provides content analysis like word count and estimated read time.

> 📦 Built with simplicity, performance, and readability in mind.

---

## ✨ Features

- ✅ Extract meta tags (title, description, keywords, author, OpenGraph)
- 🖼️ Scrape all images with `src` and `alt`
- 🔗 Collect all anchor links with their text
- 🧠 Analyze page content for:
  - Word count
  - Estimated reading time
  - Text summary (basic)
- 🌐 Minimal external dependencies

---

## 🔧 Installation

```bash
go get github.com/yarie4481/goscraper

🧪 Usage
Here's how you can use GoScraper in your project:

package main

import (
    "fmt"
    "github.com/yarie4481/goscraper/scraper"
)

func main() {
    url := "https://example.com"

    metadata, _ := scraper.GetMetadata(url)
    fmt.Println("Title:", metadata.Title)
    fmt.Println("Description:", metadata.Description)

    links, _ := scraper.GetLinks(url)
    fmt.Println("Links:", links)

    images, _ := scraper.GetImages(url)
    fmt.Println("Images:", images)

    stats, _ := scraper.AnalyzeContent(url)
    fmt.Println("Word Count:", stats.WordCount)
    fmt.Println("Read Time (min):", stats.ReadTimeMin)
    fmt.Println("Summary:", stats.Summary)
}

📁 Folder Structure
goscraper/
├── scraper/
│   ├── scraper.go       // Core logic
│   ├── helpers.go       // Internal utilities
│   └── types.go         // Reusable types
├── examples/
│   └── demo.go          // Usage demo
└── README.md

🧠 How It Works
Uses Go's net/http and golang.org/x/net/html to fetch and parse pages

Traverses HTML DOM recursively to collect metadata, links, images

Basic content analysis is done using text heuristics

No JavaScript rendering – ideal for static pages

� Types


type Metadata struct {
    Title       string
    Description string
    Keywords    []string
    OGImage     string
    Author      string
}

type Stats struct {
    WordCount   int
    ReadTimeMin int
    Summary     string
}

type Image struct {
    Src string
    Alt string
}

type Link struct {
    Text string
    Href string
}

🛠️ To-Do / Contributions Welcome
Add proxy and user-agent support

Retry on failure with backoff

Add support for sitemap parsing

Optional JavaScript rendering (with Rod or chromedp)

CLI tool

🤝 Contributing
PRs and issues are welcome. Please follow Go best practices and submit small, clear changes.

📜 License
MIT © 2025 [Yared Wubie , yaredwu02@gmail.com]

💬 Example Output


Title: Example Domain
Description: This domain is for use in illustrative examples...
Links: [Example Link - https://example.com/more]
Images: [Image{Src: "/image.jpg", Alt: "Example"}]
Word Count: 452
Read Time (min): 2
Summary: Example Domain is a placeholder page used...
```
