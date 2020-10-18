/*
 * aws-whats-new-bot
 * Copyright (c) 2020 - Puru Tuladhar (ptuladhar3@gmail.com)
 * See LICENSE file.
 */
package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

const RSS_FEED_URL = "https://aws.amazon.com/about-aws/whats-new/recent/feed/"

type WhatsNew struct {
	feed              *gofeed.Feed
	lastPublishedDate *time.Time
}

func (whatsnew *WhatsNew) getLastPublishedDate() *time.Time {
	f, err := ioutil.ReadFile("LAST_PUBLISHED_DATE")
	if err != nil {
		log.Fatalln(err)
	}
	LAST_PUBLISHED_DATE := strings.Trim(string(f), "\n")
	parsed, err := time.Parse(time.RFC1123Z, LAST_PUBLISHED_DATE)
	if err != nil {
		log.Fatalln(err)
	}
	return &parsed
}

func (whatsnew *WhatsNew) saveLastPublishedDate(lastPublishedDate *time.Time) {
	data := []byte(lastPublishedDate.Format(time.RFC1123Z))
	filename := "LAST_PUBLISHED_DATE"
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Fatalln(err)
	}
	log.Printf("Saved last published date (%s) to file: %s", data, filename)
}

func (whatsnew *WhatsNew) CheckForNewAnnouncements() bool {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(RSS_FEED_URL)
	if err != nil {
		log.Fatalln(err)
	}
	whatsnew.feed = feed

	publishedDate := feed.PublishedParsed
	log.Println("Published date:", publishedDate)

	whatsnew.lastPublishedDate = whatsnew.getLastPublishedDate()
	log.Println("Last published date:", whatsnew.lastPublishedDate)

	if publishedDate.After(*whatsnew.lastPublishedDate) {
		log.Println("📢 New announcements available!")
		whatsnew.saveLastPublishedDate(publishedDate)
		return true
	} else {
		log.Println("No announcements available!")
		return false
	}
}

func (whatsnew *WhatsNew) ListNewAnnoucements() []*gofeed.Item {
	log.Println("Listing new annoucements...")
	var annoucements []*gofeed.Item

	for _, item := range whatsnew.feed.Items {
		publishedDate := item.PublishedParsed
		title := item.Title

		if publishedDate.After(*whatsnew.lastPublishedDate) {
			log.Printf("... [%s] %s\n", publishedDate, title)
			annoucements = append(annoucements, item)
		}
	}
	return annoucements
}
