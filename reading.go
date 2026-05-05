package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

type Article struct {
	Title     string
	Link      string
	Published time.Time
	BlogName  string
}

type RSSFeed struct {
	URL  string
	Name string
}

const ARTICLE_COUNT = 5
const MAX_TITLE_LENGTH = 40

func main() {
	feeds := []RSSFeed{
		{URL: "https://mitchellh.com/feed.xml", Name: "Mitchell Hashimoto"},
		{URL: "https://erikbern.com/feed.xml", Name: "Erik Bernhardsson"},
		{URL: "https://jvns.ca/atom.xml", Name: "Julia Evans"},
		{URL: "https://feeds.feedburner.com/ThePragmaticEngineer", Name: "Gergely Orosz"},
		{URL: "https://thorstenball.com/atom.xml", Name: "Thorsten Ball"},
		{URL: "https://blog.alexellis.io/rss/", Name: "Alex Ellis"},
		{URL: "https://samwho.dev/rss.xml", Name: "Sam Rose"},
		{URL: "https://world.hey.com/dhh/feed.atom", Name: "DHH"},
		{URL: "https://research.swtch.com/feed.atom", Name: "Russ Cox"},
		{URL: "http://feeds.haacked.com/haacked", Name: "Phil Haack"},
		{URL: "https://bobheadxi.dev/feed.xml", Name: "Robert Lin"},
		{URL: "https://arslan.io/rss", Name: "Faith Arslan"},
		{URL: "https://stephango.com/feed.xml", Name: "Steph Ango"},


		{URL: "https://netflixtechblog.medium.com/feed", Name: "Netflix"},
		{URL: "https://engineering.atspotify.com/feed", Name: "Spotify"},
		{URL: "https://engineering.zalando.com/atom.xml", Name: "Zalando"},
		{URL: "https://engineering.zalando.com/atom.xml", Name: "Zalando"},
		{URL: "https://planetscale.com/blog/feed.atom", Name: "PlanetScale"},
		{URL: "https://fly.io/blog/feed.xml", Name: "Fly"},
		{URL: "https://go.dev/blog/feed.atom", Name: "Go"},

		{URL: "https://newsletter.posthog.com/feed", Name: "PostHog"},
		{URL: "https://highscalability.com/rss", Name: "High Scalability"},
	}

	var articles []Article
	parser := gofeed.NewParser()

	for _, feed := range feeds {
		log.Printf("fetching feed: %s\n", feed.Name)
		parsedFeed, err := parser.ParseURL(feed.URL)
		if err != nil {
			log.Printf("error parsing feed %s: %v\n", feed.Name, err)
			continue
		}

		count := ARTICLE_COUNT
		if len(parsedFeed.Items) < count {
			count = len(parsedFeed.Items)
		}

		for i := 0; i < count; i++ {
			item := parsedFeed.Items[i]

			// Get publication date, fallback to updated date
			var pubDate time.Time
			if item.PublishedParsed != nil {
				pubDate = *item.PublishedParsed
			} else if item.UpdatedParsed != nil {
				pubDate = *item.UpdatedParsed
			} else {
				pubDate = time.Now()
			}

			articles = append(articles, Article{
				Title:     item.Title,
				Link:      item.Link,
				Published: pubDate,
				BlogName:  feed.Name,
			})
		}
	}

	// Sort articles by publication date
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Published.After(articles[j].Published)
	})

	content := generateMarkdown(articles)

	outputDir := "src/content"
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("error creating directory: %v\n", err)
	}

	outputPath := filepath.Join(outputDir, "reading.md")
	err = os.WriteFile(outputPath, []byte(content), 0644)
	if err != nil {
		log.Fatalf("error writing file: %v\n", err)
	}

	log.Printf("successfully generated reading with %d articles\n", len(articles))
}

func generateMarkdown(articles []Article) string {
	markdown := `---
title: Reading list
description: Automatically generated from RSS feeds I follow
---

# Reading list

This page is automatically updated every night with the latest posts from blogs I follow.

`

	for _, article := range articles {
		dateStr := article.Published.Format("02.01.2006")

		title := article.Title
		if len(title) > MAX_TITLE_LENGTH {
			title = title[:MAX_TITLE_LENGTH] + "..."
		}

		markdown += fmt.Sprintf("- %s [%s](%s) - %s\n",
			dateStr, title, article.Link, article.BlogName)
	}

	return markdown
}
