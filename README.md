# 91dl
91porn视频下载器

WARNING: *小撸怡情，中撸伤身，强撸灰飞烟灭*

## 下载安装

需要go 1.16版本

```go
go get -u github.com/ilove91/91dl
```

## 使用方法

```
91porn视频下载器
WARNING: *小撸怡情，中撸伤身，强撸灰飞烟灭*

类别代码:
ori-91原创 new-最新 hot-当前最热 top-本月最热 mf-收藏最多
long-10分钟以上 longer-20分钟以上 md-本月讨论 tf-本月收藏
rf-最近加精 hd-高清 lasttop-上月最热

Usage:
  91dl [command]

Available Commands:
  help        Help about any command
  links       按照链接下载
  pages       按照页面下载
  version     打印版本号

Flags:
      --config string   配置文件 (默认不需要)
  -d, --dir string      下载到指定文件夹 (默认下载到 ./91videos/)
  -h, --help            help for 91dl
      --proxy string    网络代理, 默认支持http/socks5

Use "91dl [command] --help" for more information about a command.
```

### 按照页面下载

默认下载：“当前最热”第1页下载到当前目录下91videos/

```bash
91dl pages
```

“本月最热”第1页到第3页下载到~/mydir

```bash
91dl pages --st 1 --ed 3 -t top -d ~/mydir
```

### 按照链接下载

```bash
91dl links -d ~/mydir -v "http://91porn.com/view_video.php?viewkey=0ff9f3af6e42aab264df&page=1&viewtype=basic&category=hot,http://91porn.com/view_video.php?viewkey=71fd50381078e11ca7eb&page=1&viewtype=basic&category=hot,http://91porn.com/view_video.php?viewkey=4a2512cf4bdf9fb8abe9&page=1&viewtype=basic&category=hot"
```

### 使用代理

支持http/socks5

```bash
91dl pages --proxy http://127.0.0.1:1234
```