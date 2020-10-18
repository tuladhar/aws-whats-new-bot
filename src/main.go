/*
 * aws-whats-new-bot
 * Copyright (c) 2020 - Puru Tuladhar (ptuladhar3@gmail.com)
 * See LICENSE file.
 */
package main

import (
	"os"
	"time"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	twitter := Twitter{}
	creds := Credentials{
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
	}
	twitter.InitializeClient(&creds)
	twitter.VerifyCredentials()

	whatsnew := WhatsNew{}
	if whatsnew.CheckForNewAnnouncements() {
		announcements := whatsnew.ListNewAnnoucements()
		for _, announcement := range announcements {
			twitter.Tweet(announcement)
			time.Sleep(5 * time.Second)
		}
	}
}
