package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/swimming/go-scrapper/Job_Scrapper/parallelCSV"
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
	c := make(chan []extractedJob)

	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, len of jobs : ", len(jobs))
}

func getPage(page int, mainC chan []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

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
		go extractJob(card, c)
	})

	for i := 0; i < jobCard.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".location").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	company := cleanString(card.Find(".company").Text())
	summary := cleanString(card.Find(".summary").Text())

	c <- extractedJob{id: id,
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
		// fmt.Println(paginations, pageNum)

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

func writeJobs(jobs []extractedJob) {

	w, err := parallelCSV.NewCsvWriter("jobs.csv")
	checkErr(err)

	defer w.Flush()

	headers := []string{
		"link",
		"title",
		"location",
		"salary",
		"company",
		"summary  ",
	}

	w.Write(headers)

	writeC := make(chan bool)

	for _, job := range jobs {
		go writeJobCards(job, w, writeC)
	}

	// fmt.Println(len(jobs))
}

func writeJobCards(job extractedJob, w *parallelCSV.CsvWriter, writeC chan bool) {
	jobSlice := []string{
		baseURL + "&vjk=" + job.id,
		job.title,
		job.location,
		job.salary,
		job.company,
		job.summary,
	}

	w.Write(jobSlice)
	writeC <- true
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
