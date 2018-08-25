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
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Pallinder/go-randomdata"
	dler "github.com/joeybloggs/go-download"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func download(i int, v *video, cgn int) error {
	fmt.Printf("%3d %s\n", i, v.title)
	destFile := fmt.Sprintf("%s/%s.mp4", dir, v.title)
	if _, err := os.Stat(destFile); err == nil {
		fmt.Println("Exists, Skip")
		return nil
	}

	progress := mpb.New(
		mpb.WithWidth(60),
		mpb.WithFormat("[=>-|"),
		mpb.WithRefreshRate(150*time.Millisecond),
	)
	defer progress.Wait()
	var bs []*mpb.Bar

	options := &dler.Options{
		Request: req,
		Client:  customClient,
		Concurrency: func(size int64) int {
			if cgn == -1 {
				return gn
			}
			return cgn
		},
		Proxy: func(name string, download int, size int64, r io.Reader) io.Reader {
			bar := progress.AddBar(size,
				mpb.PrependDecorators(
					decor.CountersKibiByte("% 6.1f / % 6.1f"),
				),
				mpb.AppendDecorators(
					decor.EwmaETA(decor.ET_STYLE_MMSS, float64(size)/2048),
					decor.Name(" ] "),
					decor.AverageSpeed(decor.UnitKiB, "% .2f"),
				),
				mpb.BarRemoveOnComplete(),
			)
			bs = append(bs, bar)
			return bar.ProxyReader(r)
		},
	}

	f, err := dler.Open(v.src, options)
	if err != nil {
		for i := 0; i < len(bs); i++ {
			progress.Abort(bs[i], true)
		}
		return fmt.Errorf("fail to download %s(%s): %s", destFile, v.url, err)
	}
	defer f.Close()

	dest, err := os.Create(destFile)
	if err != nil {
		return fmt.Errorf("fail to create %s(%s): %s", destFile, v.url, err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, f)
	if err != nil {
		return fmt.Errorf("fail to save %s(%s): %s", destFile, v.url, err)
	}

	return nil
}

func req(r *http.Request) {
	r.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	r.Header.Set("Cache-Control", "max-age=0")
	r.Header.Set("Connection", "keep-alive")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	r.Header.Set("X-Forwarded-For", randomdata.IpV4Address())
}

func customClient() http.Client {
	if proxy != nil {
		return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	}
	return http.Client{}
}
