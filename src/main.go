/*
 * aws-whats-new-bot
 * Copyright (c) 2020 - Puru Tuladhar (ptuladhar3@gmail.com)
 * See LICENSE file.
 */
package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Credentials struct {
	ConsumerKey       string `envconfig:"TWITTER_CONSUMER_KEY"`
	ConsumerSecret    string `envconfig:"TWITTER_CONSUMER_SECRET"`
	AccessToken       string `envconfig:"TWITTER_ACCESS_TOKEN"`
	AccessTokenSecret string `envconfig:"TWITTER_ACCESS_TOKEN_SECRET"`
}

func main() {
	twitter := Twitter{}
	var creds Credentials
	envconfig.Process("", &creds)
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
