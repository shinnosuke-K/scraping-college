package college

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Colleges struct {
	College []struct {
		Name      string
		Pref      string
		City      string
		Station   string
		Corp      string
		Depart    string
		Deviation string
	}
}

func New() *Colleges {
	return &Colleges{}
}

func (c *Colleges) ExtractCollegeInfo(res *http.Response) error {
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find("body div#container div#contents div#main div#under div.searchResult").Each(func(i int, selection *goquery.Selection) {
		// 大学名
		collegeName := strings.TrimSpace(selection.Find("div.searchResult-list-name a").Text())
		//fmt.Println(collegeName)

		// 都道府県・市町村・駅名・（国公立or私立）
		collegeInfo := strings.Split(selection.Find("div.searchResult-list-info span.searchResult-list-profile").Text(), "/")
		address := strings.Split(collegeInfo[0], " ")

		// 学部・偏差値
		selection.Find("div.searchResult-list-gakka ul div.searchResult-list-gakubu").Each(func(i int, selection *goquery.Selection) {
			depart := strings.TrimSpace(selection.Text())
			deviation := strings.TrimSpace(selection.Next().Text())

			c.College = append(c.College, struct {
				Name      string
				Pref      string
				City      string
				Station   string
				Corp      string
				Depart    string
				Deviation string
			}{
				Name:      collegeName,
				Pref:      address[0],
				City:      address[1],
				Station:   collegeInfo[1],
				Corp:      collegeInfo[2],
				Depart:    depart,
				Deviation: deviation,
			})
		})
	})
	return nil
}
