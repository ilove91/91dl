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
	"time"

	"github.com/spf13/viper"
)

var gn int
var dir string
var proxy *url.URL
var client *http.Client

// Initialize on root
func Initialize() {
	// goroutines number
	gn = viper.GetInt("gn")
	if gn > 20 {
		gn = 20
	}
	if gn < 1 {
		gn = 5
	}

	// proxy
	proxyStr := viper.GetString("proxy")
	if proxyStr == "" {
		proxy = nil
		client = &http.Client{Timeout: time.Second * 10}
	} else {
		proxy, err := url.ParseRequestURI(proxyStr)
		if err != nil {
			log.Fatalf("Proxy Error: %s\n", err)
		}
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}, Timeout: time.Second * 10}
	}

	// saving dir
	dir = viper.GetString("dir")
	if dir == "" {
		dir = fmt.Sprintf("videos_%s", time.Now().Format("2006-01-02"))
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			log.Fatalf("Create Dir Error: %s\n", err)
		}
	}

	fmt.Println("===========================================================================")
	fmt.Printf("Saving Videos to: %s\n", dir)
	fmt.Printf("Goroutines: %d, Proxy: %s\n", gn, proxyStr)
	fmt.Println("===========================================================================")
}
