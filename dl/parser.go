// Copyright © 2018 ilove91
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dl

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/PuerkitoBio/goquery"
)

type video struct {
	url   string
	title string
	src   string
}

func getHTML(u string) (*goquery.Document, error) {
	req := buildReq(u)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func parsePage(u string) []string {
	doc, err := getHTML(u)
	if err != nil {
		log.Fatal(err)
	}

	var links []string
	doc.Find(".listchannel a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		title, _ := linkTag.Attr("title")
		if title != "" {
			links = append(links, link)
		}
	})
	return links
}

func parseVideo(u string) (*video, error) {
	for i := 0; i < 5; i++ {
		doc, err := getHTML(u)
		if err != nil {
			return nil, err
		}

		src, _ := doc.Find("video").Find("source").Attr("src")
		if src == "" {
			return nil, fmt.Errorf("no src on %s", u)
		}

		a, err := url.Parse(src)
		if err != nil {
			return nil, err
		}

		if stringNotInSlice(a.Host, excludeIPs) {
			title := doc.Find("div#viewvideo-title").Text()
			title = titleForm(title)
			return &video{u, title, src}, nil
		}
	}
	return nil, fmt.Errorf("no non-excludeIPs src on %s", u)
}

func titleForm(title string) string {
	r := strings.NewReplacer("/", " ", "\\", " ", ":", " ", "*", " ", "?", " ", "|", " ", "\"", " ", "<", " ", ">", " ")
	title = r.Replace(title)
	return strings.TrimSpace(title)
}

func buildReq(u string) *http.Request {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "91porn.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Set("X-Forwarded-For", randomdata.IpV4Address())

	return req
}

func stringNotInSlice(a string, s []string) bool {
	for _, b := range s {
		if b == a {
			return false
		}
	}
	return true
}
