package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/yen0x/karte_wiki_crawler_de/model"
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

	return k
}

func (k *KarteParser) Run() *model.Karte {
	k.FindNames()
	k.FindAttrAndRank()
	return k.Card
}

func (k *KarteParser) FindNames() {
	k.Doc.Find("big").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			k.Card.CardNameGer = s.Text()
		case 3:
			k.Card.CardNameJa = s.Text()
		case 4:
			k.Card.CardNameEn = s.Text()
		}

	})
}

func (k *KarteParser) FindAttrAndRank() {
	k.Card.MonsterAttr = k.Doc.Find("tr td i").First().Text()
	k.Card.MonsterRank = k.Doc.Find("tr td i").Get(2).FirstChild.Data[1:2]
}
