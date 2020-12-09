package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	url := os.Getenv("URL")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body div#container div#contents div#main div#under div.searchResult").Each(func(i int, selection *goquery.Selection) {
		collegeName := selection.Find("div.searchResult-list-name a").Text()
		fmt.Println(strings.TrimSpace(collegeName))

		collegeInfo := selection.Find("div.searchResult-list-info span.searchResult-list-profile").Text()
		fmt.Println(collegeInfo)

		selection.Find("div.searchResult-list-gakka ul div.searchResult-list-gakubu").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(strings.TrimSpace(selection.Text()))
			fmt.Println(strings.TrimSpace(selection.Next().Text()))
		})
	})
}
