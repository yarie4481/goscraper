package scraper

import (
	"math"
	"strings"
)

func AnalyzeContent(url string) (*Stats, error) {
	doc, err := fetchDocument(url)
	if err != nil {
		return nil, err
	}

	text := doc.Find("body").Text()
	words := strings.Fields(text)
	wordCount := len(words)
	readTime := int(math.Ceil(float64(wordCount) / 200.0)) // ~200 wpm

	summary := ""
	if len(words) > 50 {
		summary = strings.Join(words[:50], " ") + "..."
	}

	return &Stats{
		WordCount:   wordCount,
		ReadTimeMin: readTime,
		Summary:     summary,
	}, nil
}
