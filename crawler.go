package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	VNEXPRESS_URL = "https://vnexpress.net/"
)

func crawlVNExpressCategory() (map[string]string, error) {
	result := make(map[string]string)
	resp, err := getRequest(VNEXPRESS_URL)
	if err != nil {
		return nil, fmt.Errorf("error when crawling: %s", err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("error when goquery: %s", err)
	}
	doc.Find("#main_menu a").Each(func(i int, s *goquery.Selection) {
		linkRaw, _ := s.Attr("href")
		url := resolveRelative(VNEXPRESS_URL, linkRaw)
		category := s.Text()
		result[category] = url
	})
	return result, nil
}

func crawlVNExpress(category, url string) ([]Article, error) {
	var results []Article
	resp, err := getRequest(url)
	if err != nil {
		return []Article{}, fmt.Errorf("error when crawling: %s", err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return []Article{}, fmt.Errorf("error when goquery: %s", err)
	}

	doc.Find(".title_news a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		title := s.Text()
		results = append(results, Article{URL: link, Title: title, Category: category})
	})

	return results, nil
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")

	if err != nil {
		log.Print(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("error when getRequest %s", url)
	}
	return resp, nil
}

func resolveRelative(baseURL string, href string) string {
	if strings.HasPrefix(href, baseURL) {
		return href
	}

	if strings.HasPrefix(href, "/") {
		resolvedURL := fmt.Sprintf("%s%s", baseURL, href)
		return resolvedURL
	}
	return ""
}
