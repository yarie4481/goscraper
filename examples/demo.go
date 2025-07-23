package main

import (
	"fmt"

	"github.com/yarie4481/goscraper/scraper"
)

func main() {
	url := "https://example.com"

	meta, _ := scraper.GetMetadata(url)
	fmt.Println("Title:", meta.Title)
	fmt.Println("Description:", meta.Description)

	links, _ := scraper.GetLinks(url)
	fmt.Println("Links found:", len(links))

	images, _ := scraper.GetImages(url)
	fmt.Println("Images found:", len(images))

	stats, _ := scraper.AnalyzeContent(url)
	fmt.Printf("Words: %d, ReadTime: %dmin\n", stats.WordCount, stats.ReadTimeMin)
}
