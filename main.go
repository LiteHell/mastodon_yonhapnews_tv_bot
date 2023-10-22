package main

import (
	"github.com/joho/godotenv"
	"litehell.info/mastodon_yonhapnews_tv_bot/bot"
)

func main() {
	godotenv.Load()
	bot := bot.CreateBot()
	bot.Start()
	for {
	}
}
