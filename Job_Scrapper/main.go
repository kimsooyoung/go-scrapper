package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	company  string
	summary  string
}

var baseURL = "https://kr.indeed.com/jobs?q=python"

func main() {
	var jobs []extractedJob

	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		extractJobs := getPage(i)
		jobs = append(jobs, extractJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, len of jobs : ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{
		"id",
		"title",
		"location",
		"salary",
		"company",
		"summary  ",
	}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{
			baseURL + "&vjk=" + job.id,
			job.title,
			job.location,
			job.salary,
			job.company,
			job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob

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
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".location").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	company := cleanString(card.Find(".company").Text())
	summary := cleanString(card.Find(".summary").Text())

	return extractedJob{id: id,
		title:    title,
		location: location,
		salary:   salary,
		company:  company,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
