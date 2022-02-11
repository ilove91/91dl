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
	"net/http"
	"os"

	"github.com/cavaliergopher/grab/v3"
	"github.com/ilove91/91dl/utils"
	"github.com/robertkrimen/otto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	jsCode = `;var encode_version = 'jsjiami.com.v5', eexda = '__0x9ff10',  __0x9ff10=['w7FkXcKcwqs=','VMKAw7Fhw6Q=','w5nDlTY7w4A=','wqQ5w4pKwok=','dcKnwrTCtBg=','w45yHsO3woU=','54u75py15Y6177y0PcKk5L665a2j5pyo5b2156i677yg6L+S6K2D5pW65o6D5oqo5Lmn55i/5bSn5L21','RsOzwq5fGQ==','woHDiMK0w7HDiA==','54uS5pyR5Y6r7764wr3DleS+ouWtgeaesOW/sOeooe+/nei/ruitteaWsuaOmeaKiuS4o+eateW2i+S8ng==','bMOKwqA=','V8Knwpo=','csOIwoVsG1rCiUFU','5YmL6ZiV54qm5pyC5Y2i776Lw4LCrOS+muWssOacteW8lOeqtg==','w75fMA==','YsOUwpU=','wqzDtsKcw5fDvQ==','wqNMOGfCn13DmjTClg==','wozDisOlHHI=','GiPConNN','XcKzwrDCvSg=','U8K+wofCmcO6'];(function(_0x1f2e93,_0x60307d){var _0x1f9a0b=function(_0x35f19b){while(--_0x35f19b){_0x1f2e93['push'](_0x1f2e93['shift']());}};_0x1f9a0b(++_0x60307d);}(__0x9ff10,0x152));var _0x43d9=function(_0x13228a,_0x2ce452){_0x13228a=_0x13228a-0x0;var _0x424175=__0x9ff10[_0x13228a];if(_0x43d9['initialized']===undefined){(function(){var _0x270d2c=typeof window!=='undefined'?window:typeof process==='object'&&typeof require==='function'&&typeof global==='object'?global:this;var _0x58680b='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';_0x270d2c['atob']||(_0x270d2c['atob']=function(_0x5536e1){var _0x15e9d3=String(_0x5536e1)['replace'](/=+$/,'');for(var _0x4e6299=0x0,_0x3590d2,_0x48c90b,_0x557f6a=0x0,_0x2b086d='';_0x48c90b=_0x15e9d3['charAt'](_0x557f6a++);~_0x48c90b&&(_0x3590d2=_0x4e6299%0x4?_0x3590d2*0x40+_0x48c90b:_0x48c90b,_0x4e6299++%0x4)?_0x2b086d+=String['fromCharCode'](0xff&_0x3590d2>>(-0x2*_0x4e6299&0x6)):0x0){_0x48c90b=_0x58680b['indexOf'](_0x48c90b);}return _0x2b086d;});}());var _0x4a2d38=function(_0x1f120d,_0x1d6e11){var _0x4c36f9=[],_0x1c4b64=0x0,_0x18ce5c,_0x39c9fa='',_0x6d02b2='';_0x1f120d=atob(_0x1f120d);for(var _0x13b203=0x0,_0x24d88b=_0x1f120d['length'];_0x13b203<_0x24d88b;_0x13b203++){_0x6d02b2+='%'+('00'+_0x1f120d['charCodeAt'](_0x13b203)['toString'](0x10))['slice'](-0x2);}_0x1f120d=decodeURIComponent(_0x6d02b2);for(var _0x1f76f3=0x0;_0x1f76f3<0x100;_0x1f76f3++){_0x4c36f9[_0x1f76f3]=_0x1f76f3;}for(_0x1f76f3=0x0;_0x1f76f3<0x100;_0x1f76f3++){_0x1c4b64=(_0x1c4b64+_0x4c36f9[_0x1f76f3]+_0x1d6e11['charCodeAt'](_0x1f76f3%_0x1d6e11['length']))%0x100;_0x18ce5c=_0x4c36f9[_0x1f76f3];_0x4c36f9[_0x1f76f3]=_0x4c36f9[_0x1c4b64];_0x4c36f9[_0x1c4b64]=_0x18ce5c;}_0x1f76f3=0x0;_0x1c4b64=0x0;for(var _0x2b6a92=0x0;_0x2b6a92<_0x1f120d['length'];_0x2b6a92++){_0x1f76f3=(_0x1f76f3+0x1)%0x100;_0x1c4b64=(_0x1c4b64+_0x4c36f9[_0x1f76f3])%0x100;_0x18ce5c=_0x4c36f9[_0x1f76f3];_0x4c36f9[_0x1f76f3]=_0x4c36f9[_0x1c4b64];_0x4c36f9[_0x1c4b64]=_0x18ce5c;_0x39c9fa+=String['fromCharCode'](_0x1f120d['charCodeAt'](_0x2b6a92)^_0x4c36f9[(_0x4c36f9[_0x1f76f3]+_0x4c36f9[_0x1c4b64])%0x100]);}return _0x39c9fa;};_0x43d9['rc4']=_0x4a2d38;_0x43d9['data']={};_0x43d9['initialized']=!![];}var _0x302f80=_0x43d9['data'][_0x13228a];if(_0x302f80===undefined){if(_0x43d9['once']===undefined){_0x43d9['once']=!![];}_0x424175=_0x43d9['rc4'](_0x424175,_0x2ce452);_0x43d9['data'][_0x13228a]=_0x424175;}else{_0x424175=_0x302f80;}return _0x424175;};function strencode2(_0x4f0d7a){var _0x4c6de5={'Anfny':function _0x4f6a21(_0x51d0ce,_0x5a5f36){return _0x51d0ce(_0x5a5f36);}};return _0x4c6de5[_0x43d9('0x0','fo#E')](unescape,_0x4f0d7a);};(function(_0x17883e,_0x4a42d3,_0xe4080c){var _0x301ffc={'lPNHL':function _0x1c947e(_0x4d57b6,_0x51f6a5){return _0x4d57b6!==_0x51f6a5;},'EPdUx':function _0x55f4cc(_0x34b7bc,_0x9f930c){return _0x34b7bc===_0x9f930c;},'kjFfJ':'jsjiami.com.v5','DFsBH':function _0x5f08ac(_0x1e6fa1,_0x4c0aef){return _0x1e6fa1+_0x4c0aef;},'akiuH':_0x43d9('0x1','KYjt'),'VtfeI':function _0x4f3b7b(_0x572344,_0x5f0cde){return _0x572344(_0x5f0cde);},'Deqmq':_0x43d9('0x2','oYRG'),'oKQDc':_0x43d9('0x3','i^vo'),'UMyIE':_0x43d9('0x4','oYRG'),'lRwKx':function _0x5b71b4(_0x163a75,_0x4d3998){return _0x163a75===_0x4d3998;},'TOBCR':function _0x314af8(_0x3e6efe,_0x275766){return _0x3e6efe+_0x275766;},'AUOVd':_0x43d9('0x5','lALy')};_0xe4080c='al';try{if('EqF'!==_0x43d9('0x6','xSW]')){_0xe4080c+=_0x43d9('0x7','oYRG');_0x4a42d3=encode_version;if(!(_0x301ffc[_0x43d9('0x8','fo#E')](typeof _0x4a42d3,_0x43d9('0x9','*oMH'))&&_0x301ffc[_0x43d9('0xa','ov6D')](_0x4a42d3,_0x301ffc[_0x43d9('0xb','3k]D')]))){_0x17883e[_0xe4080c](_0x301ffc[_0x43d9('0xc','@&#[')]('ɾ��',_0x301ffc[_0x43d9('0xd','i^vo')]));}}else{return _0x301ffc[_0x43d9('0xe','rvlM')](unescape,input);}}catch(_0x23e6c5){if('svo'!==_0x301ffc[_0x43d9('0xf','TpCD')]){_0x17883e[_0xe4080c]('ɾ���汾�ţ�js�ᶨ�ڵ���');}else{_0xe4080c='al';try{_0xe4080c+=_0x301ffc[_0x43d9('0x10','doK*')];_0x4a42d3=encode_version;if(!(_0x301ffc[_0x43d9('0x11','ZRZ4')](typeof _0x4a42d3,_0x301ffc['UMyIE'])&&_0x301ffc[_0x43d9('0x12','@&#[')](_0x4a42d3,_0x301ffc['kjFfJ']))){_0x17883e[_0xe4080c](_0x301ffc[_0x43d9('0x13','KYjt')]('ɾ��',_0x43d9('0x14','xSW]')));}}catch(_0x4202f6){_0x17883e[_0xe4080c](_0x301ffc[_0x43d9('0x15','oYRG')]);}}}}(window));;encode_version = 'jsjiami.com.v5';`
)

var (
	destDir    string
	client     *http.Client
	grabClient *grab.Client
	jsvm       *otto.Otto
)

// Initialize on root
func Initialize() {

	// client init
	client = utils.GetNewHttpClient(10)
	grabClient = grab.NewClient()
	grabClient.HTTPClient = utils.GetNewHttpClient(0)
	grabClient.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"

	// saving dir
	destDir = viper.GetString("dir")
	if destDir == "" {
		destDir = "91videos"
	}
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, 0700); err != nil {
			log.Fatalf("Create Dir Error: %s", err)
		}
	}

	// js vm
	jsvm = otto.New()
	jsvm.Run(jsCode)

	log.Info("======================================")
	proxy := viper.GetString("proxy")
	if proxy != "" {
		log.Infof("Proxy on: %v", proxy)
	}
	log.Infof("Saving Videos to: %s", destDir)
}
