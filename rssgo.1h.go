//usr/bin/env go run $0 $@; exit

// <xbar.title>rss go/xbar.title>
// <xbar.version>v1.0</xbar.version>
// <xbar.author>senkentarou</xbar.author>
// <xbar.author.github>senkentarou</xbar.author.github>
// <xbar.desc>This is rss perser on go</xbar.desc>
// <xbar.image></xbar.image>
// <xbar.dependencies>go</xbar.dependencies>

package main

import (
  "fmt"
  "strings"
  "github.com/mmcdole/gofeed"
)

type RssUrl struct {
  title string
  url string
}

func list(fp *gofeed.Parser, url string) {
  feed, _ := fp.ParseURL(url)
  items := feed.Items

  fmt.Println("---")
  for _, item := range items{
    fmt.Println(strings.Replace(item.Title, "|", "-", -1) + " | href=" + item.Link)
  }
}

func main() {
  rssUrls := []RssUrl{
    { title: "qiita", url: "https://qiita.com/popular-items/feed.atom" },
    { title: "hatena", url: "http://b.hatena.ne.jp/hotentry/it.rss" },
  }

  fp := gofeed.NewParser()

  fmt.Println("rssgo")

  for _, rssUrl := range rssUrls{
    fmt.Println("---")
    fmt.Println(rssUrl.title)

    list(fp, rssUrl.url)
  }
}
