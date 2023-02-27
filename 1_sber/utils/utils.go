package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetPage(ctx context.Context, site string, method string,
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

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, 0, nil, fmt.Errorf("NewDocumentFromReader: %w", err)
	}

	return doc, resp.StatusCode, resp.Cookies(), nil
}

func ParseForm(doc *goquery.Document) map[string]string {
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
