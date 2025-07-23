package scraper

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
