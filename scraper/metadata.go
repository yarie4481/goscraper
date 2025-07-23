package scraper

import "strings"

func GetMetadata(url string) (*Metadata, error) {
	doc, err := fetchDocument(url)
	if err != nil {
		return nil, err
	}

	meta := &Metadata{
		Title:       doc.Find("title").Text(),
		Description: doc.Find(`meta[name="description"]`).AttrOr("content", ""),
		OGImage:     doc.Find(`meta[property="og:image"]`).AttrOr("content", ""),
		Author:      doc.Find(`meta[name="author"]`).AttrOr("content", ""),
	}

	// Keywords split
	keywords := doc.Find(`meta[name="keywords"]`).AttrOr("content", "")
	if keywords != "" {
		meta.Keywords = splitCSV(keywords)
	}

	return meta, nil
}
func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	var clean []string
	for _, p := range parts {
		trim := strings.TrimSpace(p)
		if trim != "" {
			clean = append(clean, trim)
		}
	}
	return clean
}