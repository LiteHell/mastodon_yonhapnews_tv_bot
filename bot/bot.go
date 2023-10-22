package bot

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/McKael/madon"
	"github.com/mmcdole/gofeed"
)

type Bot struct {
	mastodonBot *madon.Client
	db          *sql.DB
	itemChan    chan *gofeed.Item
}

func CreateBot() *Bot {
	var err error

	bot := new(Bot)
	bot.mastodonBot, err = createMastodonApiClient()
	if err != nil {
		panic(err)
	}
	bot.db, err = initializeDatabse(os.Getenv("DB_PATH"))
	if err != nil {
		panic(err)
	}
	bot.itemChan = make(chan *gofeed.Item)

	return bot
}

func (this Bot) Start() {
	go crawl_worker(&this)
	go post_worker(&this)
}

func crawl_worker(this *Bot) {
	for {
		// fetch feed
		log.Print("Fetching feeds...\n")
		feed, err := getFeed()
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		items := feed.Items
		log.Printf("Feed total %d items", len(items))

		// Exclude already posted news
		itemsToPost := make([]*gofeed.Item, 0)
		for _, j := range items {
			// remove unposted guids
			row := this.db.QueryRow("SELECT * from news where guid = ?", j.GUID)
			var unused string
			if row.Scan(&unused) == nil {
				log.Printf("Already posted: %s\n", j.GUID)
			} else {
				itemsToPost = append(itemsToPost, j)
			}
		}

		// Post them
		log.Printf("Total %d items to post\n", len(itemsToPost))
		for _, i := range itemsToPost {
			_, err = this.db.Exec("INSERT INTO news VALUES (?)", i.GUID)
			if err != nil {
				panic(err)
			}
			this.itemChan <- i
		}

		// wait delay
		log.Print("Waiting for delay...\n")
		time.Sleep(5000 * time.Millisecond)
	}
}

func post_worker(this *Bot) {
	for item := range this.itemChan {
		log.Printf("Posting news: %s (Title: %s)", item.GUID, item.Title)
		post, err := postNewsFeedItem(this.mastodonBot, item)

		if err != nil {
			log.Printf("Error while posting: %s", err.Error())
		} else {
			log.Printf("Post id for %s: %d", item.GUID, post.ID)
		}

		log.Printf("Waiting delay for post...")
		time.Sleep(1500 * time.Millisecond)
	}
}
