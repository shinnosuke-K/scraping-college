package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func GetPageNum(itemNum int) int {
	if itemNum <= 10 {
		return 1
	} else if itemNum%10 == 0 {
		return itemNum / 10
	} else {
		return itemNum/10 + 1
	}
}

func CheckStatus(statusCode int) error {
	if statusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d", statusCode)
	}
	return nil
}

func ExtractCollegeInfo(res *http.Response) error {
	defer res.Body.Close()
	if err := CheckStatus(res.StatusCode); err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find("body div#container div#contents div#main div#under div.searchResult").Each(func(i int, selection *goquery.Selection) {
		// 大学名
		collegeName := selection.Find("div.searchResult-list-name a").Text()
		fmt.Println(strings.TrimSpace(collegeName))

		//// 都道府県・市町村・（国公立or私立）
		//collegeInfo := selection.Find("div.searchResult-list-info span.searchResult-list-profile").Text()
		//fmt.Println(collegeInfo)
		//
		//// 学部・偏差値
		//selection.Find("div.searchResult-list-gakka ul div.searchResult-list-gakubu").Each(func(i int, selection *goquery.Selection) {
		//	fmt.Println(strings.TrimSpace(selection.Text()))
		//	fmt.Println(strings.TrimSpace(selection.Next().Text()))
		//})
	})
	return nil
}

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

	itemNum, err := strconv.Atoi(doc.Find("div.mod-pagerNum.mod-pagerNum__st span").First().Text())
	if err != nil {
		log.Fatal(err)
	}

	pageNum := GetPageNum(itemNum)
	pref := "osaka"
	for n := 1; n <= pageNum; n++ {
		url = fmt.Sprintf("https://www.minkou.jp/university/search/pref=%s/page=%d/", pref, n)

		res, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if err := ExtractCollegeInfo(res); err != nil {
			log.Fatal(err)
		}

		// 時間稼ぎ
		time.Sleep(time.Millisecond * 500)
	}
}
