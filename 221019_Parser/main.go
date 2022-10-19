package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	Debug        = false
	startLink    = "http://185.204.3.165"
	questionLink = startLink + "/question/"
	queryLimit   = 100
)

func getPage(ctx context.Context, site string, method string,
	cookies []*http.Cookie, headerValues map[string]string,
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
		return nil, 0, nil, fmt.Errorf("NewRequestWithContext error: %w", err)
	}

	if len(headerValues) > 0 {
		for header, value := range headerValues {
			req.Header.Add(header, value)
		}
	}

	if len(cookies) > 0 {
		for _, value := range cookies {
			req.AddCookie(value)
		}
	}

	if Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(string(dump))
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	var doc *goquery.Document
	if Debug {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		fmt.Println(buf.String())
		doc, err = goquery.NewDocumentFromReader(strings.NewReader(buf.String()))
	} else {
		doc, err = goquery.NewDocumentFromReader(resp.Body)
	}
	if err != nil {
		return nil, 0, nil, fmt.Errorf("NewDocumentFromReader: %w", err)
	}

	return doc, resp.StatusCode, resp.Cookies(), nil
}

func parseForm(doc *goquery.Document) map[string]string {
	formValues := make(map[string]string)

	formSelection := doc.Find("form").First()
	formSelection.Find("p").Each(func(i int, s *goquery.Selection) {
		s.Find("input").Each(func(i int, s *goquery.Selection) {
			name, _ := s.Attr("name")
			if name == "" {
				return
			}
			// process input type text
			typ, _ := s.Attr("type")
			if typ == "text" {
				formValues[name] = "test"
			}
			// process input type radio
			if typ == "radio" {
				value, _ := s.Attr("value")
				if len(value) > len(formValues[name]) {
					formValues[name] = value
				}
			}
		})
		s.Find("select").Each(func(i int, s *goquery.Selection) {
			name, _ := s.Attr("name")
			if name == "" {
				return
			}
			// process select options
			s.Find("option").Each(func(i int, s *goquery.Selection) {
				value, _ := s.Attr("value")
				if len(value) > len(formValues[name]) {
					formValues[name] = value
				}
			})
		})
	})
	return formValues
}

func main() {
	ctx := context.Background()

	// auth
	_, _, cookies, err := getPage(ctx, startLink, http.MethodGet, nil, nil, nil)
	if err != nil {
		log.Fatalf("auth error: %s", err.Error())
	}

	log.Println("auth completed")
	if Debug {
		log.Println(cookies)
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	// first request
	activeLink := questionLink + "1"
	doc, code, _, err := getPage(ctx, activeLink, http.MethodGet, cookies, headers, nil)
	if err != nil {
		log.Fatalf("get question 1 error: %s", err.Error())
	}
	if code != 200 {
		log.Fatalf("wrong question 1 answer code: %d", code)
	}

	for i := 2; i < queryLimit; i++ {
		// parse answer
		formValues := parseForm(doc)
		if len(formValues) == 0 {
			log.Fatal("parse form error")
		}
		log.Printf("question %d answer: %v", i-1, formValues)

		time.Sleep(100 * time.Millisecond) // delay for rate limits

		// next requests
		doc, code, _, err = getPage(ctx, activeLink, http.MethodPost, cookies, headers, formValues)
		if err != nil {
			log.Fatalf("get question %d error: %s", i, err.Error())
		}
		if code != 200 {
			log.Fatalf("wrong question %d answer code: %d", i, code)
		}

		// check success
		if strings.Contains(doc.Text(), "Test successfully passed") {
			log.Println("test successfully passed")
			return
		}

		activeLink = questionLink + strconv.Itoa(i)
	}
}
