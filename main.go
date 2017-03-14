package main

import (
	"fmt"
	gq "github.com/puerkitobio/goquery"
	"os"
	"regexp"
)

func main() {
	checklink := regexp.MustCompile(`archive/[[:digit:]]`)
	file, err := os.Open("test/listing.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	listing, err := gq.NewDocumentFromReader(file)
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
