package scraper

import "github.com/PuerkitoBio/goquery"

func GetImages(url string) ([]Image, error) {
	doc, err := fetchDocument(url)
	if err != nil {
		return nil, err
	}

	var images []Image
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		alt, _ := s.Attr("alt")
		if src != "" {
			images = append(images, Image{Src: src, Alt: alt})
		}
	})

	return images, nil
}
