package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yen0x/karte_wiki_crawler_de/model"
	"strings"
)

type KarteParser struct {
	Doc  *goquery.Document
	Card *model.Karte
}

func NewKarteParser(url string) *KarteParser {
	k := &KarteParser{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	} else {
		k.Doc = doc
	}
	k.Card = &model.Karte{}
	k.Card.WikiUrl = url

	return k
}

func (k *KarteParser) Run() *model.Karte {
	k.FindNames()
	k.FindAttrAndRank()
	k.FindTrAttributes()
	k.FindPictureUrl()
	k.FindCategories()
	return k.Card
}

func (k *KarteParser) FindNames() {
	k.Doc.Find("big").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			k.Card.NameGer = s.Text()
		case 3:
			k.Card.NameJa = s.Text()[3:]
		case 4:
			k.Card.NameEn = s.Text()[3:]
		}

	})
}

func (k *KarteParser) FindAttrAndRank() {
	k.Card.MonsterAttr = k.Doc.Find("tr td i").First().Text()

	if k.Doc.Find("tr td i").Length() > 2 {
		k.Card.MonsterRank = stripchars(k.Doc.Find("tr td i").Get(2).FirstChild.Data, "()")
	}
}

func (k *KarteParser) FindTrAttributes() {
	k.Doc.Find("tr td").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 21:
			k.Card.MonsterAttack = strings.TrimSpace(s.Text())
		case 22:
			k.Card.MonsterDefense = strings.TrimSpace(s.Text())
		case 23:
			k.Card.Code = strings.TrimSpace(s.Text())
		case 24:
			k.Card.EffectType = strings.TrimSpace(s.Text())
		case 26:
			k.Card.Description = strings.TrimSpace(s.Text())
		}
	})
}

func (k *KarteParser) FindPictureUrl() {
	k.Card.PictureUrl, _ = k.Doc.Find("tr td a img").First().Attr("src")
	k.Card.PictureUrl = BASE_URL + k.Card.PictureUrl
}

func (k *KarteParser) FindCategories() {
	k.Doc.Find(".mw-normal-catlinks ul li a").Each(func(i int, s *goquery.Selection) {
		k.Card.Categories = append(k.Card.Categories, s.Text())
	})
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
