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
	//url := "http://yugioh-wiki.de/wiki/Kategorie:Yugioh_Karte"
	url := "http://yugioh-wiki.de/w/index.php?title=Kategorie:Yugioh_Karte&pagefrom=Zweiklauen-Angriff#mw-pages"

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
			go func() {
				if len(nextPage) > 0 {
					fmt.Println("NEXT PAGE " + nextPage)
					lparser.FindCardLinks(nextPage)
					nextPageChannel <- lparser.FindNextLinkPage(nextPage)
				}
			}()
		case link := <-lparser.LinksChannel:
			go func() {
				for retry := 0; retry < 2; retry++ {
					kparser := parser.NewKarteParser(link)
					if kparser != nil {
						card := kparser.Run()
						//			data, _ := json.Marshal(card)
						fmt.Printf("Parsing card %d\n", linkCounter)
						linkCounter++
						//			fmt.Println(string(data))
						cards = append(cards, card)
					}
				}
			}()
		default:
			if previousCounter == linkCounter {
				if duration := time.Since(previousTime); duration.Seconds() > 5.0 {
					fmt.Println("Quitting")
					parse = false
				}
			} else {
				previousTime = time.Now()
				previousCounter = linkCounter
			}

		}
	}

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

func getCards(c chan *model.Karte, url string) {
	fmt.Printf("---%s---\n", url)
	kparser := parser.NewKarteParser(url)
	c <- kparser.Run()
}

func printer(card *model.Karte) {
	fmt.Println("PRINTER")
	data, _ := json.Marshal(card)
	fmt.Println(string(data))
}

func x() {
	if len(nextPage) > 0 {
		fmt.Println("NEXT PAGE " + nextPage)
	}
	lparser.FindCardLinks(nextPage)
	nextPageChannel <- lparser.FindNextLinkPage(nextPage)
}
