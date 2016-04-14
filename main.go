package main

import (
	"encoding/json"
	"fmt"
	"github.com/yen0x/karte_wiki_crawler_de/parser"
)

func main() {
	lparser := parser.NewCardLinkParser("http://yugioh-wiki.de/wiki/Kategorie:Yugioh_Karte")
	links, nextPage := lparser.Run()
	linkCounter := 0
	for len(nextPage) > 0 {
		fmt.Println("NEXT PAGE: " + nextPage)

		for _, url := range links {
			fmt.Printf("%d -- %s\n", linkCounter, url)
			linkCounter++
			kparser := parser.NewKarteParser(url)
			card := kparser.Run()
			data, _ := json.Marshal(card)
			fmt.Println(string(data))
		}
		lparser.SetUrl(nextPage)
		links, nextPage = lparser.Run()
	}
}
