package douban

import (
	"os"
	"strings"

	"github.com/wheat-os/slubby/stream"
)

type DoubanPipline struct {
	f *os.File
}

func (t *DoubanPipline) OpenSpider() error {
	fs, err := os.Create("./data.csv")
	if err != nil {
		return err
	}
	t.f = fs
	fs.WriteString("name,author,desc,commit,score\n")
	return err
}

func (t *DoubanPipline) CloseSpider() error {
	t.f.Close()
	return nil
}

func (t *DoubanPipline) ProcessItem(item stream.Item) stream.Item {
	filmItem := item.(*DoubanItem)
	t.f.WriteString(strings.Join(append([]string(nil), filmItem.Name, filmItem.Author,
		filmItem.Desc, filmItem.Commit, filmItem.Score), ","))
	t.f.WriteString("\n")
	return item
}
