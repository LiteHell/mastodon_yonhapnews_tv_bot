package bot

import (
	"os"

	"github.com/McKael/madon"
)

func createMastodonApiClient() (mc *madon.Client, err error) {
	appId := os.Getenv("CLIENT_ID")
	appSecret := os.Getenv("CLIENT_SECRET")
	userToken := madon.UserToken{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
	}

	return madon.RestoreApp("연합뉴스TV 봇", "social.litehell.info", appId, appSecret, &userToken)
}
