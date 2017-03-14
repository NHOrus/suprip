package main

import (
	"fmt"
	gq "github.com/puerkitobio/goquery"
	"os"
	"regexp"
)

func main() {
	riptext()
}

func riptext() {
	file, err := os.Open("test/page.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	text, err := gq.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	ArtistName := "Artemis"
	text.Find(".postContainer").Each(func(i int, s *gq.Selection) {
		if s.Find(".desktop .name").Text() != ArtistName {
			return
		}
		fmt.Println(s.Text())
	})
}

func riplinks() {
	checklink := regexp.MustCompile(`archive/[[:digit:]]`)
	file, err := os.Open("test/listing.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	listing, err := gq.NewDocumentFromReader(file)
	//listing, err := gq.NewDocument("http://suptg.thisisnotatrueending.com/archive.html?tags=eclipsed+moon")
	if err != nil {
		fmt.Println(err)
		return
	}
	table := listing.Find("table")
	table.Find("a").Each(func(i int, s *gq.Selection) {
		if val, ok := s.Attr("href"); ok {
			if checklink.MatchString(val) {
				fmt.Println(val)
			}
		}
	})
}
