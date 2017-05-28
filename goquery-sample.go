package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(word string) {
	doc, err := goquery.NewDocument("http://qiita.com/search?q=" + word)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".searchResult").Each(func(i int, s *goquery.Selection) {
		a := s.Find("h1 > a")
		title := a.Text()
		path, _ := a.Attr("href")
		url := "http://qiita.com" + path
		fmt.Printf("%d : %s  : %s \n", i+1, title, url)
	})
}

func main() {
	ExampleScrape("golang")
}
