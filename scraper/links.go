package scraper

import "github.com/PuerkitoBio/goquery"

func GetLinks(url string) ([]Link, error) {
	doc, err := fetchDocument(url)
	if err != nil {
		return nil, err
	}

	var links []Link
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		href, _ := s.Attr("href")
		if href != "" {
			links = append(links, Link{Text: text, Href: href})
		}
	})

	return links, nil
}
