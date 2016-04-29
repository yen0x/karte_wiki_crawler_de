package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type LinkParser struct {
	Doc          *goquery.Document
	LinksChannel chan string
}

func (l *LinkParser) SetChannel(c chan string) {
	l.LinksChannel = c
}

func (l *LinkParser) FindCardLinks(url string) (links []string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error:" + err.Error())
		return
	}
	doc.Find("div#mw-pages div.mw-category div.mw-category-group ul li a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		l.LinksChannel <- BASE_URL + href
	})
	return
}

func (l *LinkParser) FindNextLinkPage(url string) (href string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	doc.Find("div#mw-pages > a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "nächste Seite" {
			href, _ = s.Attr("href")
			href = BASE_URL + href
			return
		}
	})
	return
}
