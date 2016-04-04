package main

import (
	"encoding/json"
	"fmt"
	"github.com/yen0x/karte_wiki_crawler_de/parser"
)

func main() {
	kparser := parser.NewKarteParser("http://yugioh-wiki.de/wiki/Chaos_End-Meister")

	card := kparser.Run()
	data, _ := json.Marshal(card)
	fmt.Println(string(data))
}
