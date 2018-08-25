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
)

var baseURL = "http://91porn.com"

// LinksDl download by links
func LinksDl(vlinks []string) {
	var vs []*video
	for i := 0; i < len(vlinks); i++ {
		v, err := parseVideo(vlinks[i])
		if err != nil {
			fmt.Printf("Cannot parse url to video: %s\n", err)
			continue
		}
		fmt.Println(v.title)
		vs = append(vs, v)
	}

	fmt.Printf("Find %d Videos\n", len(vs))
	var wvs []*video
	for i := 0; i < len(vs); i++ {
		if err := download(i, vs[i], -1); err != nil {
			fmt.Println(err)
			wvs = append(wvs, vs[i])
		}
	}

	fmt.Printf("Redownloading %d Error Videos\n", len(wvs))
	for i := 0; i < len(wvs); i++ {
		if err := download(i, wvs[i], 1); err != nil {
			fmt.Println(err)
		}
	}
}

// PagesDl download by pages
// category: new hot rp long md tf mf rf top top-1 hd
func PagesDl(p1 int, p2 int, t string) {
	if p1 > p2 {
		p1 = p2
	}
	fmt.Printf("Download category %s from page %d to %d\n", t, p1, p2)
	fmt.Println("===========================================================================")

	var url string
	for i := p1; i <= p2; i++ {
		switch t {
		case "new":
			url = fmt.Sprintf("%s/v.php?next=watch&page=%d", baseURL, i)
		case "top-1":
			url = fmt.Sprintf("%s/v.php?category=%s&m=-1&viewtype=basic&page=%d", baseURL, t, i)
		default:
			url = fmt.Sprintf("%s/v.php?category=%s&viewtype=basic&page=%d", baseURL, t, i)
		}
		vl := parseList(url)
		fmt.Printf("Downloading page %d ...\n", i)
		LinksDl(vl)
		fmt.Println("===========================================================================")
	}
}
