package bot

import (
	"fmt"
	"log"
	"mime"
	"os"

	"github.com/McKael/madon"
	"github.com/mmcdole/gofeed"
)

func postNewsFeedItem(mastodon *madon.Client, feedItem *gofeed.Item) (status *madon.Status, err error) {
	hasThumbnail := len(feedItem.Enclosures) != 0
	medias := []int64{}
	if hasThumbnail {
		// Assign first enclosure
		enclosure := feedItem.Enclosures[0]

		// Create temp file
		ext, _ := mime.ExtensionsByType(enclosure.Type)
		tmpFile, _ := os.CreateTemp("", fmt.Sprintf("*%s", ext[0]))

		// Download to temporary file
		downloadFile(tmpFile, enclosure.URL)
		fileName := tmpFile.Name()
		tmpFile.Close()

		// Upload media
		attachment, err := mastodon.UploadMedia(fileName, fmt.Sprintf("썸네일 이미지 (%s)", feedItem.Title), "")

		// Delete temporary file
		os.Remove(tmpFile.Name())

		// Use uploaded media if there's no error
		if err == nil {
			medias = []int64{attachment.ID}
		} else {
			log.Printf("Error while uploading media: %s", err.Error())
		}
	}

	// Post news
	return mastodon.PostStatus(
		fmt.Sprintf("%s\n%s", feedItem.Title, feedItem.Link), // feedItem.Description,
		0,
		medias,
		false,
		"", //fmt.Sprintf("%s\n%s", feedItem.Title, feedItem.Link),
		"public",
	)
}
