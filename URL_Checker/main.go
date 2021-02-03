package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var urlErr = errors.New("URL Get Failed!!")

func hitURLNormal(url string) error {
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return urlErr
	}
	return nil
}

func hitURLGoRoutine(url string, c chan<- requestResult) {
	res, err := http.Get(url)

	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAIL"
	}
	c <- requestResult{url: url, status: status}
}

func withoutGoRoutine(urls []string) {
	results := make(map[string]string)

	fmt.Println("Waiting for URL Get...")

	for _, url := range urls {
		result := "OK"
		err := hitURLNormal(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func awesomeGoRoutine(urls []string) {
	results := make(map[string]string)
	c := make(chan requestResult)

	for _, url := range urls {
		go hitURLGoRoutine(url, c)
	}

	fmt.Println("Waiting for URL Get...")

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func main() {

	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.netflix.com",
		"https://www.soundcloud.com",
		"https://www.notion.com",
		"https://www.helloworld.com",
	}

	withoutGoRoutine(urls)
	awesomeGoRoutine(urls)
}
