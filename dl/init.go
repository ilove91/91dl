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
	"os"

	"github.com/cavaliercoder/grab"
	"github.com/spf13/viper"
)

var dir string
var proxy *url.URL
var client *http.Client
var dler *grab.Client

// Initialize on root
func Initialize() {
	// proxy
	proxyStr := viper.GetString("proxy")
	if proxyStr == "" {
		proxy = nil
		client = &http.Client{}
	} else {
		proxy, err := url.ParseRequestURI(proxyStr)
		if err != nil {
			log.Fatalf("Proxy Error: %s\n", err)
		}
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	}
	dler = grab.NewClient()
	dler.HTTPClient = client

	// saving dir
	dir = viper.GetString("dir")
	if dir == "" {
		dir = "91videos"
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			log.Fatalf("Create Dir Error: %s\n", err)
		}
	}

	fmt.Println("===========================================================================")
	fmt.Printf("Saving Videos to: %s\n", dir)
	if proxyStr != "" {
		fmt.Printf("Proxy on: %s\n", proxyStr)
	}
	fmt.Println("===========================================================================")
}
