package spider

import (
	"net/url"
	"testing"
	"time"
)

func TestCrawler(t *testing.T) {
	r := &Requireds{
		SiteUrl:    &url.URL{Scheme: "https", Host: "github.com"},
		TimeOuT:    10,
		MaxDepth:   2,
		Concurrent: 5,
		Delay:      1000,
	}

	crawler, err := r.NewCrawler()
	if err != nil {
		t.Fatalf("Failed to create crawler: %v", err)
	}

	// Start the crawler
	go crawler.Start(true)

	// Wait for some time to let the crawler do its work
	time.Sleep(10 * time.Second)

	// Check results
	if crawler.UrlSet.Size() == 0 {
		t.Error("Expected URLs to be collected")
	}

	if crawler.JsSet.Size() == 0 {
		t.Error("Expected JavaScript URLs to be collected")
	}
}
