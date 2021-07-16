<p align="center"><img src="static/logo.png" alt="Annie" height="100px"></p>

<div align="center">
  <a href="https://codecov.io/gh/iawia002/annie">
    <img src="https://img.shields.io/codecov/c/github/iawia002/annie.svg?style=flat-square" alt="Codecov">
  </a>
  <a href="https://github.com/iawia002/annie/actions">
    <img src="https://img.shields.io/github/workflow/status/iawia002/annie/ci?style=flat-square" alt="GitHub Workflow Status">
  </a>
  <a href="https://goreportcard.com/report/github.com/iawia002/annie">
    <img src="https://goreportcard.com/badge/github.com/iawia002/annie?style=flat-square" alt="Go Report Card">
  </a>
  <a href="https://github.com/iawia002/annie/releases">
    <img src="https://img.shields.io/github/release/iawia002/annie.svg?style=flat-square" alt="GitHub release">
  </a>
  <a href="https://formulae.brew.sh/formula/annie">
    <img src="https://img.shields.io/homebrew/v/annie.svg?style=flat-square" alt="Homebrew">
  </a>
  <a href="https://t.me/anniedev">
    <img src="https://img.shields.io/badge/telegram-join%20chat-0088cc.svg?longCache=true&style=flat-square" alt="telegram">
  </a>
</div>


👾 Annie is a fast, simple and clean video downloader built with Go.

- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Install via `go get`](#install-via-go-get)
  - [Homebrew (macOS only)](#homebrew-macos-only)
  - [Arch Linux](#arch-linux)
  - [Void Linux](#void-linux)
  - [Scoop on Windows](#scoop-on-windows)
  - [Chocolatey on Windows](#chocolatey-on-windows)
- [Getting Started](#getting-started)
  - [Download a video](#download-a-video)
  - [Download anything else](#download-anything-else)
  - [Download playlist](#download-playlist)
  - [Multiple inputs](#multiple-inputs)
  - [Resume a download](#resume-a-download)
  - [Auto retry](#auto-retry)
  - [Cookies](#cookies)
  - [Proxy](#proxy)
  - [Multi-Thread](#multi-thread)
  - [Short link](#short-link)
    - [bilibili](#bilibili)
  - [Use specified Referrer](#use-specified-referrer)
  - [Specify the output path and name](#specify-the-output-path-and-name)
  - [Debug Mode](#debug-mode)
  - [Reuse extracted data](#reuse-extracted-data)
  - [Options](#options)
    - [Download:](#download)
    - [Network:](#network)
    - [Playlist:](#playlist)
    - [Filesystem:](#filesystem)
    - [Subtitle:](#subtitle)
    - [Youku:](#youku)
    - [aria2:](#aria2)
- [Supported Sites](#supported-sites)
- [Known issues](#known-issues)
  - [优酷](#优酷)
- [Contributing](#contributing)
- [Authors](#authors)
- [Similar projects](#similar-projects)
- [License](#license)


## Installation

### Prerequisites

The following dependencies are required and must be installed separately.

* **[FFmpeg](https://www.ffmpeg.org)**

> **Note**: FFmpeg does not affect the download, only affects the final file merge.

### Install via `go get`

To install Annie, use `go get`, or download the binary file from [Releases](https://github.com/iawia002/annie/releases) page.

```bash
$ GO111MODULE=on go get github.com/iawia002/annie
```

### Homebrew (macOS only)

For macOS users, you can install `annie` via:

```bash
$ brew install annie
```

### Arch Linux

For Arch Users [AUR](https://aur.archlinux.org/packages/annie) package is available.

### Void Linux

For Void linux users, you can install `annie` via:

```
$ xbps-install -S annie
```

### [Scoop](https://scoop.sh/) on Windows

```sh
$ scoop install annie
```

### [Chocolatey](https://chocolatey.org/) on Windows

```
$ choco install annie
```

## Getting Started

Usage:

```
annie [OPTIONS] URL [URL...]
```

### Download a video

```console
$ annie https://www.youtube.com/watch?v=dQw4w9WgXcQ

 Site:      YouTube youtube.com
 Title:     Rick Astley - Never Gonna Give You Up (Video)
 Type:      video
 Stream:
     [248]  -------------------
     Quality:         1080p video/webm; codecs="vp9"
     Size:            63.93 MiB (67038963 Bytes)
     # download with: annie -f 248 ...

 41.88 MiB / 63.93 MiB [=================>-------------]  65.51% 4.22 MiB/s 00m05s
```

> Note: wrap the URL in quotation marks if it contains special characters. (thanks @tonyxyl for pointing this out)
>
> `$ annie 'https://...'`

The `-i` option displays all available quality of video without downloading.

```console
$ annie -i https://www.youtube.com/watch?v=dQw4w9WgXcQ

 Site:      YouTube youtube.com
 Title:     Rick Astley - Never Gonna Give You Up (Video)
 Type:      video
 Streams:   # All available quality
     [248]  -------------------
     Quality:         1080p video/webm; codecs="vp9"
     Size:            49.29 MiB (51687554 Bytes)
     # download with: annie -f 248 ...

     [137]  -------------------
     Quality:         1080p video/mp4; codecs="avc1.640028"
     Size:            43.45 MiB (45564306 Bytes)
     # download with: annie -f 137 ...

     [398]  -------------------
     Quality:         720p video/mp4; codecs="av01.0.05M.08"
     Size:            37.12 MiB (38926432 Bytes)
     # download with: annie -f 398 ...

     [136]  -------------------
     Quality:         720p video/mp4; codecs="avc1.4d401f"
     Size:            31.34 MiB (32867324 Bytes)
     # download with: annie -f 136 ...

     [247]  -------------------
     Quality:         720p video/webm; codecs="vp9"
     Size:            31.03 MiB (32536181 Bytes)
     # download with: annie -f 247 ...
```

Use `annie -f stream "URL"` to download a specific stream listed in the output of `-i` option.

### Download anything else

If Annie is provided the URL of a specific resource, then it will be downloaded directly:

```console
$ annie https://img9.bcyimg.com/drawer/15294/post/1799t/1f5a87801a0711e898b12b640777720f.jpg

annie doesn't support this URL right now, but it will try to download it directly

 Site:      Universal
 Title:     1f5a87801a0711e898b12b640777720f
 Type:      image/jpeg
 Stream:
     [default]  -------------------
     Size:            1.00 MiB (1051042 Bytes)
     # download with: annie -f default "URL"

 1.00 MiB / 1.00 MiB [===================================] 100.00% 1.21 MiB/s 0s
```

### Download playlist

The `-p` option downloads an entire playlist instead of a single video.

```console
$ annie -i -p https://www.bilibili.com/bangumi/play/ep198061

 Site:      哔哩哔哩 bilibili.com
 Title:     Doctor X 第四季：第一集
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            845.66 MiB (886738354 Bytes)
     # download with: annie -f default "URL"


 Site:      哔哩哔哩 bilibili.com
 Title:     Doctor X 第四季：第二集
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            930.71 MiB (975919195 Bytes)
     # download with: annie -f default "URL"

......
```

You can use the `-start`, `-end` or `-items` option to specify the download range of the list:

```
-start
    	Playlist video to start at (default 1)
-end
    	Playlist video to end at
-items
    	Playlist video items to download. Separated by commas like: 1,5,6,8-10
```

For bilibili playlists only:

```
-eto
  File name of each bilibili episode doesn't include the playlist title
```

### Multiple inputs

You can also download multiple URLs at once:

```console
$ annie https://www.bilibili.com/video/av21877586 https://www.bilibili.com/video/av21990740

 Site:      哔哩哔哩 bilibili.com
 Title:     【莓机会了】甜到虐哭的13集单集MAD「我现在什么都不想干,更不想看14集」
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            51.88 MiB (54403767 Bytes)
     # download with: annie -f default "URL"


 Site:      哔哩哔哩 bilibili.com
 Title:     【莓救了】甜到虐哭！！！国家队单集MAD-当熟悉的bgm响起，眼泪从脸颊滑下
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            77.63 MiB (81404093 Bytes)
     # download with: annie -f default "URL"
```

These URLs will be downloaded one by one.

You can also use the `-F` option to read URLs from file:

```console
$ annie -F ~/Desktop/u.txt

 Site:      微博 weibo.com
 Title:     在Google，我们设计什么？ via@阑夕
 Type:      video
 Stream:
     [default]  -------------------
     Size:            19.19 MiB (20118196 Bytes)
     # download with: annie -f default "URL"

 19.19 MiB / 19.19 MiB [=================================] 100.00% 9.69 MiB/s 1s

......
```

You can use the `-start`, `-end` or `-items` option to specify the download range of the list:

```
-start
    	File line to start at (default 1)
-end
    	File line to end at
-items
    	File lines to download. Separated by commas like: 1,5,6,8-10
```

### Resume a download

<kbd>Ctrl</kbd>+<kbd>C</kbd> interrupts a download.

A temporary `.download` file is kept in the output directory. If `annie` is ran with the same arguments, then the download progress will resume from the last session.

### Auto retry

annie will auto retry when the download failed, you can specify the retry times by `-retry` option (default is 100).

### Cookies

Cookies can be provided to `annie` with the `-c` option if they are required for accessing the video.

Cookies can be the following format or [Netscape Cookie](https://curl.haxx.se/rfc/cookie_spec.html) format:

```console
name=value; name2=value2; ...
```

Cookies can be a string or a text file, supply cookies in one of the two following ways.

As a string:

```console
$ annie -c "name=value; name2=value2" https://www.bilibili.com/video/av20203945
```

As a text file:

```console
$ annie -c cookies.txt https://www.bilibili.com/video/av20203945
```

If the `-c` is not set, `annie` will try to get the cookies from the current user's Chrome or Edge automatically.
To use this feature, you need to shutdown your Chrome or Edge for only one time and let `annie` launch the browser for you.

### Proxy

You can set the HTTP/SOCKS5 proxy using environment variables:

```console
$ HTTP_PROXY="http://127.0.0.1:1087/" annie -i https://www.youtube.com/watch?v=Gnbch2osEeo
```

```console
$ HTTP_PROXY="socks5://127.0.0.1:1080/" annie -i https://www.youtube.com/watch?v=Gnbch2osEeo
```

### Multi-Thread

Use `-n` option to set the number of download threads(default is 10, only works for multiple-parts video).

> **Special Tips:** Use too many threads in **mgtv** download will cause HTTP 403 error, we recommend setting the number of threads to **1**.

### Short link
#### bilibili

You can just use `av` or `ep` number to download bilibili's video:

```console
$ annie -i ep198381 av21877586

 Site:      哔哩哔哩 bilibili.com
 Title:     狐妖小红娘：第79话 南国公主的吃货本色
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            485.23 MiB (508798478 Bytes)
     # download with: annie -f default "URL"


 Site:      哔哩哔哩 bilibili.com
 Title:     【莓机会了】甜到虐哭的13集单集MAD「我现在什么都不想干,更不想看14集」
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            51.88 MiB (54403767 Bytes)
     # download with: annie -f default "URL"
```

### Use specified Referrer

A Referrer can be used for the request with the `-r` option:


```console
$ annie -r https://www.bilibili.com/video/av20383055/ http://cn-scnc1-dx.acgvideo.com/...

...
```

### Specify the output path and name

The `-o` option sets the path, and `-O` option sets the name of the downloaded file:

```console
$ annie -o ../ -O "hello" https://...
```

### Debug Mode

The `-d` option outputs network request messages:

```console
$ annie -i -d http://www.bilibili.com/video/av20088587

URL:         http://www.bilibili.com/video/av20088587
Method:      GET
Headers:     http.Header{
    "Referer":         {"http://www.bilibili.com/video/av20088587"},
    "Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
    "Accept-Charset":  {"UTF-8,*;q=0.5"},
    "Accept-Encoding": {"gzip,deflate,sdch"},
    "Accept-Language": {"en-US,en;q=0.8"},
    "User-Agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.146 Safari/537.36"},
}
Status Code: 200

URL:         https://interface.bilibili.com/v2/playurl?appkey=84956560bc028eb7&cid=32782944&otype=json&qn=116&quality=116&type=&sign=fb2e3f261fec398652f96d358517e535
Method:      GET
Headers:     http.Header{
    "Accept-Charset":  {"UTF-8,*;q=0.5"},
    "Accept-Encoding": {"gzip,deflate,sdch"},
    "Accept-Language": {"en-US,en;q=0.8"},
    "User-Agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.146 Safari/537.36"},
    "Referer":         {"https://interface.bilibili.com/v2/playurl?appkey=84956560bc028eb7&cid=32782944&otype=json&qn=116&quality=116&type=&sign=fb2e3f261fec398652f96d358517e535"},
    "Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
}
Status Code: 200

 Site:      哔哩哔哩 bilibili.com
 Title:     燃油动力的遥控奥迪R8跑赛道
 Type:      video
 Streams:   # All available quality
     [default]  -------------------
     Quality:         高清 1080P
     Size:            64.38 MiB (67504795 Bytes)
     # download with: annie -f default "URL"
```

### Reuse extracted data

The `-j` option will print the extracted data in JSON format.

```console
$ annie -j https://www.bilibili.com/video/av20203945

{
    "site": "哔哩哔哩 bilibili.com",
    "title": "【2018拜年祭单品】相遇day by day",
    "type": "video",
    "streams": {
        "15": {
            "urls": [
                {
                    "url": "...",
                    "size": 18355205,
                    "ext": "flv"
                }
            ],
            "quality": "流畅 360P",
            "size": 18355205
        },
        "32": {
            "urls": [
                {
                    "url": "...",
                    "size": 40058632,
                    "ext": "flv"
                }
            ],
            "quality": "清晰 480P",
            "size": 40058632
        },
        "64": {
            "urls": [
                {
                    "url": "...",
                    "size": 82691087,
                    "ext": "flv"
                }
            ],
            "quality": "高清 720P",
            "size": 82691087
        },
        "80": {
            "urls": [
                {
                    "url": "...",
                    "size": 121735559,
                    "ext": "flv"
                }
            ],
            "quality": "高清 1080P",
            "size": 121735559
        }
    }
}
```

### Options

```
  -i	Information only
  -F string
    	URLs file path
  -d	Debug mode
  -j	Print extracted data
  -v	Show version
```

#### Download:

```
  -f string
    	Select specific stream to download
  -p	Download playlist
  -n int
    	The number of download thread (only works for multiple-parts video) (default 10)
  -c string
    	Cookie
  -r string
    	Use specified Referrer
  -cs int
    	HTTP chunk size for downloading (in MB) (default 0)
```

#### Network:

```
  -retry int
    	How many times to retry when the download failed (default 10)
```

#### Playlist:

```
  -start int
    	Playlist video to start at (default 1)
  -end int
    	Playlist video to end at
  -items string
    	Playlist video items to download. Separated by commas like: 1,5,6,8-10
```

#### Filesystem:

```
  -o string
    	Specify the output path
  -O string
    	Specify the output file name
```

#### Subtitle:

```
  -C	Download captions
```

#### Youku:

```
  -ccode string
    	Youku ccode (default "0590")
  -ckey string
    	Youku ckey (default "7B19C0AB12633B22E7FE81271162026020570708D6CC189E4924503C49D243A0DE6CD84A766832C2C99898FC5ED31F3709BB3CDD82C96492E721BDD381735026")
  -password string
    	Youku password
```

#### aria2:

> Note: If you use aria2 to download, you need to merge the multi-part videos yourself.

```
  -aria2
    	Use Aria2 RPC to download
  -aria2addr string
    	Aria2 Address (default "localhost:6800")
  -aria2method string
    	Aria2 Method (default "http")
  -aria2token string
    	Aria2 RPC Token
```


## Supported Sites

| Site       | URL                          | 🎬 Videos | 🌁 Images | 📚 Playlist | 🍪 VIP adaptation |
| ---------- | ---------------------------- | -------- | -------- | ---------- | ---------------- |
| 抖音       | <https://www.douyin.com>     | ✓        |          |            |                  |
| 哔哩哔哩   | <https://www.bilibili.com>   | ✓        |          | ✓          | ✓                |
| 半次元     | <https://bcy.net>            |          | ✓        |            |                  |
| pixivision | <https://www.pixivision.net> |          | ✓        |            |                  |
| 优酷       | <https://www.youku.com>      | ✓        |          |            | ✓                |
| YouTube    | <https://www.youtube.com>    | ✓        |          | ✓          |                  |
| 爱奇艺     | <https://www.iqiyi.com>      | ✓        |          |            |                  |
| 芒果TV     | <https://www.mgtv.com>       | ✓        |          |            |                  |
| 糖豆广场舞 | <http://www.tangdou.com>     | ✓        |          | ✓          |                  |
| Tumblr     | <https://www.tumblr.com>     | ✓        | ✓        |            |                  |
| Vimeo      | <https://vimeo.com>          | ✓        |          |            |                  |
| Facebook   | <https://facebook.com>       | ✓        |          |            |                  |
| 斗鱼视频   | <https://v.douyu.com>        | ✓        |          |            |                  |
| 秒拍       | <https://www.miaopai.com>    | ✓        |          |            |                  |
| 微博       | <https://weibo.com>          | ✓        |          |            |                  |
| Instagram  | <https://www.instagram.com>  | ✓        | ✓        |            |                  |
| Twitter    | <https://twitter.com>        | ✓        |          |            |                  |
| 腾讯视频   | <https://v.qq.com>           | ✓        |          |            |                  |
| 网易云音乐 | <https://music.163.com>      | ✓        |          |            |                  |
| 音悦台     | <https://yinyuetai.com>      | ✓        |          |            |                  |
| 极客时间   | <https://time.geekbang.org>  | ✓        |          |            |                  |
| Pornhub    | <https://pornhub.com>        | ✓        |          |            |                  |
| XVIDEOS    | <https://xvideos.com>        | ✓        |          |            |                  |
| 聯合新聞網 | <https://udn.com>            | ✓        |          |            |                  |
| TikTok     | <https://www.tiktok.com>     | ✓        |          |            |                  |
| 好看视频   | <https://haokan.baidu.com>   | ✓        |          |            |                  |
| AcFun      | <https://www.acfun.cn>       | ✓        |          | ✓          |                  |
| Eporner    | <https://eporner.com>        | ✓        |          |            |                  |


## Known issues

### 优酷

优酷的 `ccode` 经常变化导致 annie 不可用，如果你知道有新的可用的 `ccode`，可以直接使用 `annie -ccode ...` 而不用等待 annie 更新（当然，也欢迎你给我们提一个 Pull request 来更新默认的 `ccode`）

最好是每次下载都附带登录过的 Cookie 以避免部分 `ccode` 的问题


## Contributing

Annie is an open source project and built on the top of open-source projects. If you are interested, then you are welcome to contribute. Let's make Annie better, together. 💪

Check out the [Contributing Guide](./CONTRIBUTING.md) to get started.

Special thanks to [@Yasujizr](https://github.com/Yasujizr) who designed the amazing logo!

Thanks for [JetBrains](https://www.jetbrains.com/?from=annie) for the wonderful IDE.

<a href="https://www.jetbrains.com/?from=annie"><img src="static/jetbrains-variant-3.svg" /></a>

## Authors

Code with ❤️ by [iawia002](https://github.com/iawia002) and lovely [contributors](https://github.com/iawia002/annie/graphs/contributors)


## Similar projects

* [youtube-dl](https://github.com/rg3/youtube-dl)
* [you-get](https://github.com/soimort/you-get)
* [ytdl](https://github.com/rylio/ytdl)


## License

MIT

Copyright (c) 2018-present, iawia002
