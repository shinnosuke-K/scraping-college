package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/shinnosuke-K/scraping-college/college"

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

	c := college.New()
	for n := 1; n <= pageNum; n++ {
		url = fmt.Sprintf("https://www.minkou.jp/university/search/pref=%s/page=%d/", pref, n)

		res, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if err := c.ExtractCollegeInfo(res); err != nil {
			log.Fatal(err)
		}

		// 時間稼ぎ
		rand.Seed(time.Now().Unix())
		time.Sleep(time.Millisecond * time.Duration(100*rand.Intn(10)))
	}
}
