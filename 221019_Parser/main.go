package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getPage(ctx context.Context, site string, method string, cookies []*http.Cookie,
	formValues map[string]string) (*goquery.Document, int, []*http.Cookie, error) {

	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}

	body := io.Reader(nil)
	if len(formValues) > 0 {
		data := url.Values{}
		for key, value := range formValues {
			data.Add(key, value)
		}
		body = strings.NewReader(data.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, site, body)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	if len(cookies) > 0 {
		for _, value := range cookies {
			req.AddCookie(value)
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("NewDocumentFromReader: %w", err)
	}

	return doc, resp.StatusCode, resp.Cookies(), nil
}

func parseForm(doc *goquery.Document) {
	doc.Find("form").Each(func(_ int, s *goquery.Selection) {
		action, _ := s.Attr("action")
		method, _ := s.Attr("method")
		fmt.Println(action, method)
		s.Find("p").Each(func(_ int, s *goquery.Selection) {
			s.Find("input").Each(func(_ int, s *goquery.Selection) {
				name, _ := s.Attr("name")
				if name == "" {
					return
				}
				typ, _ := s.Attr("type")
				if typ == "text" {
					fmt.Println(name, ":text")
				}
				if typ == "radio" {
					fmt.Println(name, ":radio")
				}
			})
			s.Find("select").Each(func(_ int, s *goquery.Selection) {
				name, _ := s.Attr("name")
				if name == "" {
					return
				}
				fmt.Println(name, ":select")
			})
		})
	})
}

func main() {
	ctx := context.Background()
	_, _, cookies, err := getPage(ctx, "http://185.204.3.165", http.MethodGet, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cookies)
	doc, _, cookies, err := getPage(ctx, "http://185.204.3.165/question/1", http.MethodGet, cookies, nil)
	if err != nil {
		log.Fatal(err)
	}
	parseForm(doc)

	// cookies, err := getCookies()

	// Load the HTML document
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	//
	// }

	// Find the review items
	// doc.Find(".question-summary .summary").Each(func(i int, s *goquery.Selection) {
	// 	title := s.Find("H3").Text()
	// 	fmt.Println(i, title)
	// })
}
