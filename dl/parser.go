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
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type video struct {
	webURL    string
	title     string
	videoSrc  string
	mediaType string
}

func getHTML(u string) (*goquery.Document, error) {
	req := buildReq(u)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
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
	doc.Find(".videos-text-align a").Each(func(index int, item *goquery.Selection) {
		link, _ := item.Attr("href")
		log.Infof("%3d url: %v", index+1, link)
		if strings.HasPrefix(link, "http://91porn.com/view_video.php?viewkey=") {
			links = append(links, link)
		}
	})
	return links
}

func parseVideo(u string) (*video, error) {
	doc, err := getHTML(u)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	title = strings.ReplaceAll(title, ".", " ")
	title = strings.ReplaceAll(title, "Chinese homemade video", "")
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "\n", "")

	author := doc.Find(".title-yakov a span").Text()
	author = strings.TrimSpace(author)

	encrypted := doc.Find("video").Find("script").Text()
	if encrypted == "" {
		return nil, fmt.Errorf("found no encrypt str")
	}
	compile := regexp.MustCompile(`document.write\(strencode2\("(.*)"`)
	submatch := compile.FindAllStringSubmatch(encrypted, -1)
	if len(submatch) != 1 {
		return nil, fmt.Errorf("parse encrypt str err, encrypt str: %s, submatch: %s", encrypted, submatch)
	}
	encrypted = submatch[0][1]

	decrypted, err := jsvm.Call("strencode2", nil, encrypted)
	if err != nil {
		return nil, fmt.Errorf("js decrypt err: %v, encrypt str: %v", err, encrypted)
	}

	compile = regexp.MustCompile(`<source src='(.*)' type=`)
	submatch = compile.FindAllStringSubmatch(decrypted.String(), -1)
	videoSrc := submatch[0][1]
	if _, err := url.Parse(videoSrc); err != nil {
		return nil, fmt.Errorf("video src %s", videoSrc)
	}

	vNumber, mediaType, err := parseMedia(videoSrc)
	if err != nil {
		return nil, err
	}
	title = fmt.Sprintf("[%s] [%s] %s", vNumber, author, title)
	r := strings.NewReplacer("\\", " ", "/", " ", ":", " ", "*", " ", "?", " ", "\"", " ", "<", " ", ">", " ", "|", " ")
	title = r.Replace(title)

	return &video{u, title, videoSrc, mediaType}, nil
}

func parseMedia(src string) (vNumber, mediaType string, err error) {
	// m3u8
	compile := regexp.MustCompile(`m3u8/[\d]*/([\d]*).m3u8`)
	submatch := compile.FindAllStringSubmatch(src, -1)
	if len(submatch) == 1 {
		vNumber := submatch[0][1]
		return vNumber, "m3u8", nil
	}

	// mp4
	compile = regexp.MustCompile(`mp43/([\d]*).mp4\?st=.*`)
	submatch = compile.FindAllStringSubmatch(src, -1)
	if len(submatch) == 1 {
		vNumber := submatch[0][1]
		return vNumber, "mp4", nil
	}

	return "", "", fmt.Errorf("no media found in video src: %v", src)
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
