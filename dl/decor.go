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
	"github.com/vbauerster/mpb/decor"
)

func customETA(style int, resp *grab.Response, wcc ...decor.WC) decor.Decorator {
	var wc decor.WC
	for _, widthConf := range wcc {
		wc = widthConf
	}
	wc.Init()
	d := &myETA{
		WC:    wc,
		style: style,
		resp:  resp,
	}
	return d
}

type myETA struct {
	decor.WC
	style       int
	completeMsg *string
	resp        *grab.Response
}

func (d *myETA) Decor(st *decor.Statistics) string {
	if st.Completed && d.completeMsg != nil {
		return d.FormatMsg(*d.completeMsg)
	}

	var str string

	remaining := d.resp.ETA().Sub(time.Now())
	hours := int64((remaining / time.Hour) % 60)
	minutes := int64((remaining / time.Minute) % 60)
	seconds := int64((remaining / time.Second) % 60)

	switch d.style {
	case decor.ET_STYLE_GO:
		str = fmt.Sprint(time.Duration(remaining.Seconds()) * time.Second)
	case decor.ET_STYLE_HHMMSS:
		str = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	case decor.ET_STYLE_HHMM:
		str = fmt.Sprintf("%02d:%02d", hours, minutes)
	case decor.ET_STYLE_MMSS:
		str = fmt.Sprintf("%02d:%02d", minutes, seconds)
	}

	return d.FormatMsg(str)
}

func (d *myETA) OnCompleteMessage(msg string) {
	d.completeMsg = &msg
}

func customSpeed(unit int, unitFormat string, resp *grab.Response, wcc ...decor.WC) decor.Decorator {
	var wc decor.WC
	for _, widthConf := range wcc {
		wc = widthConf
	}
	wc.Init()
	d := &mySpeed{
		WC:         wc,
		unit:       unit,
		unitFormat: unitFormat,
		resp:       resp,
	}
	return d
}

type mySpeed struct {
	decor.WC
	unit        int
	unitFormat  string
	msg         string
	completeMsg *string
	resp        *grab.Response
}

func (d *mySpeed) Decor(st *decor.Statistics) string {
	if st.Completed {
		if d.completeMsg != nil {
			return d.FormatMsg(*d.completeMsg)
		}
		return d.FormatMsg(d.msg)
	}

	speed := d.resp.BytesPerSecond()

	switch d.unit {
	case decor.UnitKiB:
		d.msg = fmt.Sprintf(d.unitFormat, decor.SpeedKiB(speed))
	case decor.UnitKB:
		d.msg = fmt.Sprintf(d.unitFormat, decor.SpeedKB(speed))
	default:
		d.msg = fmt.Sprintf(d.unitFormat, speed)
	}

	return d.FormatMsg(d.msg)
}

func (d *mySpeed) OnCompleteMessage(msg string) {
	d.completeMsg = &msg
}
