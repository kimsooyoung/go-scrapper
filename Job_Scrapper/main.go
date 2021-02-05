package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL = "https://kr.indeed.com/jobs?q=python"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println(pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCard := doc.Find(".jobsearch-SerpJobCard")
	jobCard.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := card.Find(".title>a").Text()
		location := card.Find(".sjcl>a").Text()
		fmt.Println(id, title, location)
	})
}

func getPages() int {
	pageNum := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// Find the review items
	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		paginations := s.Find("li").Text()
		pageNum = s.Find("li").Length()
		fmt.Println(paginations, pageNum)

		curPage := string(paginations[0])

		if curPage == "1" {
			pageNum -= 1
		} else if pageNum < 7 {
			pageNum -= 1
		} else {
			pageNum -= 2
		}
	})

	return pageNum
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln()
	}
}

func checkRes(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Status Code Error")
	}
}
