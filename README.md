# 91dl
91porn downloader &amp; spider

## Installation

```go
go get -u github.com/ilove91/91dl
```

## Usage

```
Downloader for 91porn

Category: new-最新 hot-当前最热 rp-最近得分 long-10分钟以上
          md-本月讨论 tf-本月收藏 mf-收藏最多 rf-最近加精
          top-本月最热 top-1-上月最热 hd-高清

Usage:
  91dl [command]

Available Commands:
  help        Help about any command
  links       Download videos by links
  pages       Download videos by pages with category
  version     Print version info

Flags:
      --config string   config file (default is ./config.yaml)
  -d, --dir string      directory to save videos (default is ./videos_'date'/)
      --gn int          number of goroutines to download (default 5)
  -h, --help            help for 91dl
      --proxy string    net proxy, support http/socks5

Use "91dl [command] --help" for more information about a command.
```
