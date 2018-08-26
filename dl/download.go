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
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func download(i int, v *video, cgn int) error {
	fmt.Printf("Grab %3d %s\n", i, v.title)

	// build request
	req, _ := grab.NewRequest("", v.src)
	req.Filename = fmt.Sprintf("%s/%s.mp4", dir, v.title)

	// start download
	resp := dler.Do(req)

	// progress bar
	p := mpb.New(
		mpb.WithWidth(60),
		mpb.WithFormat("[=>-|"),
		mpb.WithRefreshRate(200*time.Millisecond),
	)
	defer p.Wait()

	bar := p.AddBar(resp.Size,
		mpb.PrependDecorators(
			decor.Elapsed(decor.ET_STYLE_MMSS),
			decor.CountersKibiByte(" % 6.1f / % 6.1f"),
		),
		mpb.AppendDecorators(
			customETA(decor.ET_STYLE_MMSS, resp),
			decor.Name(" ] "),
			customSpeed(decor.UnitKiB, "% .2f", resp),
		),
		mpb.BarRemoveOnComplete(),
	)

	complete := 0
	delta := 0

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			delta = int(resp.BytesComplete()) - complete
			bar.IncrBy(delta)
			complete += delta
		case <-resp.Done:
			delta = int(resp.BytesComplete()) - complete
			bar.IncrBy(delta)
			complete += delta
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		p.Abort(bar, true)
		return fmt.Errorf("download error %s: %s", req.Filename, err)
	}
	return nil
}
