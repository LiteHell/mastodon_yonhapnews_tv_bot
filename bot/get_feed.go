package bot

import "github.com/mmcdole/gofeed"

func getFeed() (feed *gofeed.Feed, err error) {
	url := "http://www.yonhapnewstv.co.kr/browse/feed/"

	parser := gofeed.NewParser()
	return parser.ParseURL(url)
}
