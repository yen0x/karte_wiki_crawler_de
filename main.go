package main

import (
	"encoding/json"
	"fmt"
	"github.com/yen0x/karte_wiki_crawler_de/model"
	"github.com/yen0x/karte_wiki_crawler_de/parser"
	"io/ioutil"
	"time"
)

func main() {
	lparser := parser.LinkParser{}
	lparser.LinksChannel = make(chan string)

	nextPageChannel := make(chan string)
	cards := make([]*model.Karte, 0)
	url := "http://yugioh-wiki.de/wiki/Kategorie:Yugioh_Karte"
	//url := "http://yugioh-wiki.de/w/index.php?title=Kategorie:Yugioh_Karte&pagefrom=Zweiklauen-Angriff#mw-pages"

	linkCounter := 0
	previousCounter := 0
	previousTime := time.Now()
	parse := true

	go func() {
		nextPageChannel <- url
	}()
	for parse {
		select {
		case nextPage := <-nextPageChannel:
			go getLinks(nextPage, lparser, nextPageChannel)
		case link := <-lparser.LinksChannel:
			go getCards(link, &linkCounter, &cards)
		default:
			quitParsing(&previousCounter, &linkCounter, &previousTime, &parse)
		}
	}
	// Print to file
	cardsBytes, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("Printing output file")
	err = ioutil.WriteFile("output.json", cardsBytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File written")

	/*var input string
	fmt.Scanln(&input)*/
}

func getLinks(nextPage string, lparser parser.LinkParser, nextPageChannel chan string) {
	if len(nextPage) > 0 {
		fmt.Println("Extracting cards from " + nextPage)
	}
	lparser.FindCardLinks(nextPage)
	nextLinkPage := lparser.FindNextLinkPage(nextPage)
	if nextLinkPage != "" {
		nextPageChannel <- nextLinkPage
	}
}

func getCards(link string, linkCounter *int, cards *[]*model.Karte) {
	for retry := 0; retry < 2; retry++ {
		kparser := parser.NewKarteParser(link)
		if kparser != nil {
			card := kparser.Run()
			//			data, _ := json.Marshal(card)
			fmt.Printf("Parsing card %d\n", *linkCounter)
			*linkCounter++
			//			fmt.Println(string(data))
			*cards = append(*cards, card)
		}
	}
}

func quitParsing(previousCounter, linkCounter *int, previousTime *time.Time, parse *bool) {
	if *previousCounter == *linkCounter {
		if duration := time.Since(*previousTime); duration.Seconds() > 5.0 {
			fmt.Println("Quitting")
			*parse = false
		}
	} else {
		*previousTime = time.Now()
		*previousCounter = *linkCounter
	}
}
