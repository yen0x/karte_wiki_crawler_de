package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type LinkParser struct {
	Doc *goquery.Document
}

func NewCardLinkParser(url string) *LinkParser {
	l := &LinkParser{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	} else {
		l.Doc = doc
	}
	return l
}

func (l *LinkParser) SetUrl(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	} else {
		l.Doc = doc
	}
}

func (l *LinkParser) Run() ([]string, string) {
	links := l.FindCardLinks()
	next := l.FindNextLinkPage()
	return links, next
}

func (l *LinkParser) FindCardLinks() (links []string) {
	l.Doc.Find("div#mw-pages div.mw-category div.mw-category-group ul li a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		links = append(links, BASE_URL+href)
	})
	return
}

func (l *LinkParser) FindNextLinkPage() (href string) {
	l.Doc.Find("div#mw-pages > a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "nächste Seite" {
			href, _ = s.Attr("href")
			href = BASE_URL + href
			return
		}
	})
	return
}
