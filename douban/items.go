package douban

import "github.com/wheat-os/slubby/stream"

type DoubanItem struct {
	stream.Item
	Name   string `csv:"名称"`
	Author string `csv:"author"`
	Desc   string `csv:"desc"`
	Commit string `csv:"commit"`
	Score  string `csv:"score"`
}
