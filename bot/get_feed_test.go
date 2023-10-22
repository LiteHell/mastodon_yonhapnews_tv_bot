package bot

import "testing"

func TestGetFeed(t *testing.T) {
	feed, err := getFeed()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if feed.Len() == 0 {
		t.Fatal("No items on feed")
	}
	firstItem := feed.Items[0]
	t.Logf("GUID: %s\nTitle: %s\nPublished on: %v\nEnclosure url:%s\nLink: %s\nContent: %s", firstItem.GUID, firstItem.Title, firstItem.PublishedParsed, firstItem.Enclosures[0].URL, firstItem.Link, firstItem.Content)
	//feed.Items[0].UpdatedParsed
	//t.Logf("%+v", feed)
}
