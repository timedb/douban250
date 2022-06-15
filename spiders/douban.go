package spiders

import (
	"douban/douban"
	"fmt"
	"strings"
	"sync"

	"github.com/wheat-os/slubby/spider"
	"github.com/wheat-os/slubby/stream"
	"github.com/wheat-os/wlog"
)

type doubanSpider struct{}

func (t *doubanSpider) UId() string {
	return "douban"
}

func (t *doubanSpider) FQDN() string {
	return "movie.douban.com"
}

func (t *doubanSpider) Parse(response *stream.HttpResponse) (stream.Stream, error) {
	lis := response.Xpath(`//ol[@class="grid_view"]/li`)

	stms := make([]stream.Stream, 0, len(lis))
	for _, li := range lis {
		hd := li.Xpath(`.//div[@class="hd"]/a/span/text()`)
		name := hd[0].Data

		body := li.Xpath(`.//div[@class="bd"]/p/text()`)
		author := strings.TrimSpace(body[0].Data)
		desc := strings.TrimSpace(body[1].Data)

		start := li.Xpath(`.//div[@class="star"]//span/text()`)
		score := start[0].Data
		commit := start[1].Data

		stms = append(stms, &douban.DoubanItem{
			Item:   stream.BasicItem(t),
			Name:   name,
			Author: author,
			Desc:   desc,
			Commit: commit,
			Score:  score,
		})
	}
	return stream.StreamLists(t, stms...), nil
}

func (t *doubanSpider) StartRequest() stream.Stream {
	stms, err := stream.StreamListRangeInt(t, func(i int) (stream.Stream, error) {
		url := "https://movie.douban.com/top250?start=%d"
		url = fmt.Sprintf(url, i*25)
		return stream.Request(t, url, nil)
	}, 0, 9)

	if err != nil {
		wlog.Panic(err)
	}
	return stms
}

var (
	doubanOnce = sync.Once{}
	doubanP    *doubanSpider
)

func DoubanSpider() spider.Spider {
	doubanOnce.Do(func() {
		doubanP = &doubanSpider{}
	})
	return doubanP
}
